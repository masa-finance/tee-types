package job

// WebScraperType represents the job type for web scraping
const WebScraperType = "webscraper"

// WebScraperConfiguration defines configuration for web scraping
type WebScraperConfiguration struct {
	Blacklist []string `json:"blacklist"`
}

// WebScraperArgs defines arguments for web scraping jobs
type WebScraperArgs struct {
	URL      string `json:"url"`
	Selector string `json:"selector"`
	MaxDepth int    `json:"max_depth"`
}
