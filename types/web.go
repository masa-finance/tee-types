package types

type WebSearchArguments struct {
	URL      string `json:"url"`
	Selector string `json:"selector"`
	MaxDepth int    `json:"max_depth"`
}
