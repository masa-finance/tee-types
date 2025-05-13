package jobs

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

// Section represents a selected section of a web page
type Section struct {
	Text     string `json:"text"`
	HTML     string `json:"html"`
	Selector string `json:"selector"`
}

// CollectedData represents data collected from web scraping
type CollectedData struct {
	URL      string    `json:"url"`
	Title    string    `json:"title"`
	Sections []Section `json:"sections"`
}
