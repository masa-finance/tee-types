package jobs

import (
	"time"
)

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

// TweetResult represents a Tweet returned from Twitter
type TweetResult struct {
	ID             int64  `json:"id"`
	TweetID        string `json:"tweet_id"`
	ConversationID string `json:"conversation_id"`
	UserID         string `json:"user_id"`
	Text           string `json:"text"`
	CreatedAt      time.Time `json:"created_at"`
	Timestamp      int64  `json:"timestamp"`
	
	IsQuoted     bool     `json:"is_quoted"`
	IsPin        bool     `json:"is_pin"`
	IsReply      bool     `json:"is_reply"`
	IsRetweet    bool     `json:"is_retweet"`
	IsSelfThread bool     `json:"is_self_thread"`
	Likes        int      `json:"likes"`
	Hashtags     []string `json:"hashtags"`
	HTML         string   `json:"html"`
	Replies      int      `json:"replies"`
	Retweets     int      `json:"retweets"`
	URLs         []string `json:"urls"`
	Username     string   `json:"username"`
	
	Photos []Photo `json:"photos"`
	Videos []Video `json:"videos"`
	
	RetweetedStatusID string `json:"retweeted_status_id"`
	Views             int    `json:"views"`
	SensitiveContent  bool   `json:"sensitive_content"`
	
	// Fields from TwitterX API
	AuthorID          string       `json:"author_id"`
	PublicMetrics     PublicMetrics `json:"public_metrics"`
	PossiblySensitive bool         `json:"possibly_sensitive"`
	Lang              string       `json:"lang"`
	NewestID          string       `json:"newest_id"`
	OldestID          string       `json:"oldest_id"`
	ResultCount       int          `json:"result_count"`
	
	Error error `json:"-"`
}

// PublicMetrics represents public metrics for a Tweet
type PublicMetrics struct {
	RetweetCount    int `json:"retweet_count"`
	ReplyCount      int `json:"reply_count"`
	LikeCount       int `json:"like_count"`
	QuoteCount      int `json:"quote_count"`
	BookmarkCount   int `json:"bookmark_count"`
	ImpressionCount int `json:"impression_count"`
}

// Photo represents an image attached to a Tweet
type Photo struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

// Video represents a video attached to a Tweet
type Video struct {
	ID      string `json:"id"`
	Preview string `json:"preview"`
	URL     string `json:"url"`
	HLSURL  string `json:"hls_url"`
}
