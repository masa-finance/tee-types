package types

type WebSearchArguments struct {
	URL      string `json:"url"`
	Selector string `json:"selector"`
	Depth    int    `json:"depth"`
	MaxDepth int    `json:"max_depth"`
}
