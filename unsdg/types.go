package unsdg

// Goal is a UN Sustainable Development Goal.
type Goal struct {
	Rank        int    `json:"rank"`
	Code        string `json:"code"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

// Target is a specific target under a UN SDG goal.
type Target struct {
	Rank        int    `json:"rank"`
	Code        string `json:"code"`
	Goal        string `json:"goal"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// wire types match the raw API response shapes.
type wireGoal struct {
	Code        string `json:"code"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URI         string `json:"uri"`
}

type wireTarget struct {
	Goal        string `json:"goal"`
	Code        string `json:"code"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URI         string `json:"uri"`
}
