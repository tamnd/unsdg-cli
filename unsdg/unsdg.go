// Package unsdg is the library behind the unsdg command.
package unsdg

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

// DefaultUserAgent identifies the client to the UN SDG API.
const DefaultUserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"

// Config holds constructor parameters.
type Config struct {
	BaseURL   string
	UserAgent string
	Rate      time.Duration
	Retries   int
	Timeout   time.Duration
}

// DefaultConfig returns sensible defaults.
func DefaultConfig() Config {
	return Config{
		BaseURL:   "https://unstats.un.org/sdgapi",
		UserAgent: DefaultUserAgent,
		Rate:      300 * time.Millisecond,
		Retries:   3,
		Timeout:   30 * time.Second,
	}
}

// Client talks to the UN SDG API.
type Client struct {
	cfg        Config
	httpClient *http.Client
	mu         sync.Mutex
	last       time.Time
}

// NewClient returns a Client with the given config.
func NewClient(cfg Config) *Client {
	return &Client{
		cfg:        cfg,
		httpClient: &http.Client{Timeout: cfg.Timeout},
	}
}

func (c *Client) get(ctx context.Context, url string) ([]byte, error) {
	var lastErr error
	for attempt := 0; attempt <= c.cfg.Retries; attempt++ {
		if attempt > 0 {
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-time.After(backoff(attempt)):
			}
		}
		b, retry, err := c.do(ctx, url)
		if err == nil {
			return b, nil
		}
		lastErr = err
		if !retry {
			return nil, err
		}
	}
	return nil, fmt.Errorf("get: %w", lastErr)
}

func (c *Client) do(ctx context.Context, url string) ([]byte, bool, error) {
	c.pace()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, false, err
	}
	req.Header.Set("User-Agent", c.cfg.UserAgent)
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, true, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode == http.StatusTooManyRequests || resp.StatusCode >= 500 {
		return nil, true, fmt.Errorf("http %d", resp.StatusCode)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, false, fmt.Errorf("http %d", resp.StatusCode)
	}

	b, err := io.ReadAll(io.LimitReader(resp.Body, 8<<20))
	if err != nil {
		return nil, true, err
	}
	return b, false, nil
}

func (c *Client) pace() {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.cfg.Rate <= 0 {
		return
	}
	if wait := c.cfg.Rate - time.Since(c.last); wait > 0 {
		time.Sleep(wait)
	}
	c.last = time.Now()
}

func backoff(attempt int) time.Duration {
	d := time.Duration(attempt) * 500 * time.Millisecond
	if d > 5*time.Second {
		d = 5 * time.Second
	}
	return d
}

// Goals fetches all 17 SDGs.
func (c *Client) Goals(ctx context.Context, limit int) ([]Goal, error) {
	url := c.cfg.BaseURL + "/v1/sdg/Goal/List"
	raw, err := c.get(ctx, url)
	if err != nil {
		return nil, err
	}
	var wires []wireGoal
	if err := json.Unmarshal(raw, &wires); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}
	goals := make([]Goal, 0, len(wires))
	for i, w := range wires {
		if limit > 0 && i >= limit {
			break
		}
		goals = append(goals, wireToGoal(w, i+1))
	}
	return goals, nil
}

// Targets fetches SDG targets. If goal is empty, fetches all targets.
// If goal is a number like "1", fetches only targets for that goal.
func (c *Client) Targets(ctx context.Context, goal string, limit int) ([]Target, error) {
	var url string
	if goal == "" {
		url = c.cfg.BaseURL + "/v1/sdg/Target/List"
	} else {
		url = c.cfg.BaseURL + "/v1/sdg/Target/List?goal=" + goal
	}
	raw, err := c.get(ctx, url)
	if err != nil {
		return nil, err
	}
	var wires []wireTarget
	if err := json.Unmarshal(raw, &wires); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}
	targets := make([]Target, 0, len(wires))
	for i, w := range wires {
		if limit > 0 && i >= limit {
			break
		}
		targets = append(targets, wireToTarget(w, i+1))
	}
	return targets, nil
}

func wireToGoal(w wireGoal, rank int) Goal {
	url := "https://sdgs.un.org/goals/goal" + w.Code
	desc := truncate(w.Description, 120)
	return Goal{
		Rank:        rank,
		Code:        w.Code,
		Title:       w.Title,
		Description: desc,
		URL:         url,
	}
}

func wireToTarget(w wireTarget, rank int) Target {
	desc := truncate(w.Description, 120)
	return Target{
		Rank:        rank,
		Code:        w.Code,
		Goal:        w.Goal,
		Title:       w.Title,
		Description: desc,
	}
}

func truncate(s string, n int) string {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "\n", " ")
	rs := []rune(s)
	if len(rs) <= n {
		return s
	}
	return string(rs[:n-1]) + "..."
}
