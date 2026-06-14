package unsdg_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tamnd/unsdg-cli/unsdg"
)

func TestGoals(t *testing.T) {
	payload, _ := json.Marshal([]map[string]string{
		{"code": "1", "title": "No Poverty", "description": "End poverty in all its forms everywhere.", "uri": "/v1/sdg/Goal/1"},
		{"code": "2", "title": "Zero Hunger", "description": "End hunger, achieve food security.", "uri": "/v1/sdg/Goal/2"},
	})

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("User-Agent") == "" {
			t.Error("request carried no User-Agent")
		}
		_, _ = w.Write(payload)
	}))
	defer srv.Close()

	cfg := unsdg.DefaultConfig()
	cfg.BaseURL = srv.URL
	cfg.Rate = 0

	c := unsdg.NewClient(cfg)
	goals, err := c.Goals(context.Background(), 10)
	if err != nil {
		t.Fatal(err)
	}
	if len(goals) != 2 {
		t.Fatalf("got %d goals, want 2", len(goals))
	}
	if goals[0].Title != "No Poverty" {
		t.Errorf("title = %q, want No Poverty", goals[0].Title)
	}
	if goals[0].Rank != 1 {
		t.Errorf("rank = %d, want 1", goals[0].Rank)
	}
}

func TestGoalsLimit(t *testing.T) {
	items := make([]map[string]string, 5)
	for i := range items {
		items[i] = map[string]string{"code": "1", "title": "Goal", "description": "desc", "uri": "/v1/sdg/Goal/1"}
	}
	payload, _ := json.Marshal(items)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write(payload)
	}))
	defer srv.Close()

	cfg := unsdg.DefaultConfig()
	cfg.BaseURL = srv.URL
	cfg.Rate = 0

	c := unsdg.NewClient(cfg)
	goals, err := c.Goals(context.Background(), 3)
	if err != nil {
		t.Fatal(err)
	}
	if len(goals) != 3 {
		t.Fatalf("got %d goals, want 3 (limit applied)", len(goals))
	}
}

func TestTargets(t *testing.T) {
	payload, _ := json.Marshal([]map[string]string{
		{"goal": "1", "code": "1.1", "title": "Eradicate extreme poverty", "description": "By 2030, eradicate extreme poverty.", "uri": "/v1/sdg/Target/1.1"},
		{"goal": "1", "code": "1.2", "title": "Reduce poverty by half", "description": "By 2030, reduce at least by half.", "uri": "/v1/sdg/Target/1.2"},
	})

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write(payload)
	}))
	defer srv.Close()

	cfg := unsdg.DefaultConfig()
	cfg.BaseURL = srv.URL
	cfg.Rate = 0

	c := unsdg.NewClient(cfg)
	targets, err := c.Targets(context.Background(), "1", 10)
	if err != nil {
		t.Fatal(err)
	}
	if len(targets) != 2 {
		t.Fatalf("got %d targets, want 2", len(targets))
	}
	if targets[0].Code != "1.1" {
		t.Errorf("code = %q, want 1.1", targets[0].Code)
	}
}
