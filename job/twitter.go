package job

// TwitterJobTypes defines the various types of Twitter jobs
const (
	// TwitterScraperType represents standard Twitter scraping jobs
	TwitterScraperType = "twitter"
	
	// TwitterCredentialScraperType represents Twitter scraping jobs using credentials
	TwitterCredentialScraperType = "twitter-credential"
	
	// TwitterApiScraperType represents Twitter scraping jobs using API keys
	TwitterApiScraperType = "twitter-api"
)

// TwitterKeyAuthType defines the type of authentication for Twitter API keys
type TwitterKeyAuthType string

// TwitterKeyAuthType constants
const (
	CredentialAuthType TwitterKeyAuthType = "credential"
	ApiKeyAuthType     TwitterKeyAuthType = "apikey"
	UnknownAuthType    TwitterKeyAuthType = "unknown"
)

// TwitterApiKeyType defines the type of Twitter API key
type TwitterApiKeyType string

// TwitterApiKeyType constants
const (
	TwitterApiKeyTypeBase       TwitterApiKeyType = "base"
	TwitterApiKeyTypeElevated   TwitterApiKeyType = "elevated"
	TwitterApiKeyTypeCredential TwitterApiKeyType = "credential"
	TwitterApiKeyTypeUnknown    TwitterApiKeyType = "unknown"
)

// TwitterScraperConfiguration defines configuration for Twitter scraping
type TwitterScraperConfiguration struct {
	Accounts              []string `json:"twitter_accounts"`
	ApiKeys               []string `json:"twitter_api_keys"`
	DataDir               string   `json:"data_dir"`
	SkipLoginVerification bool     `json:"skip_login_verification,omitempty"` // If true, skips Twitter's verify_credentials check
}

// TwitterScraperArgs defines arguments for Twitter scraping jobs
type TwitterScraperArgs struct {
	SearchType string `json:"type"`
	Query      string `json:"query"`
	Count      int    `json:"count"`
	MaxResults int    `json:"max_results"`
	NextCursor string `json:"next_cursor"`
}
