// Package types provides shared types between tee-worker and tee-indexer
package types

import "time"

// TwitterSearchArguments defines arguments for Twitter searches
type TwitterSearchArguments struct {
	QueryType  string `json:"type"`  // Optional, type of search
	Query      string `json:"query"` // Username or search query
	Count      int    `json:"count"`
	StartTime  string `json:"start_time"`  // Optional ISO timestamp
	EndTime    string `json:"end_time"`    // Optional ISO timestamp
	MaxResults int    `json:"max_results"` // Optional, max number of results
	NextCursor string `json:"next_cursor"`
}

type TweetResult struct {
	ID             int64 `json:"id"`
	TweetID        string
	ConversationID string
	UserID         string
	Text           string
	CreatedAt      time.Time
	Timestamp      int64

	ThreadCursor struct {
		FocalTweetID string
		ThreadID     string
		Cursor       string
		CursorType   string
	}
	IsQuoted     bool
	IsPin        bool
	IsReply      bool
	IsRetweet    bool
	IsSelfThread bool
	Likes        int
	Hashtags     []string
	HTML         string
	Replies      int
	Retweets     int
	URLs         []string
	Username     string

	Photos []Photo

	// Video type.
	Videos []Video

	RetweetedStatusID string
	Views             int
	SensitiveContent  bool

	// from twitterx
	AuthorID          string
	PublicMetrics     PublicMetrics
	PossiblySensitive bool
	Lang              string
	NewestID          string
	OldestID          string
	ResultCount       int

	Error error
}

type PublicMetrics struct {
	RetweetCount    int
	ReplyCount      int
	LikeCount       int
	QuoteCount      int
	BookmarkCount   int
	ImpressionCount int
}
type Photo struct {
	ID  string
	URL string
}

type Video struct {
	ID      string
	Preview string
	URL     string
	HLSURL  string
}
