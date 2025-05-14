// Package types provides shared types between tee-worker and tee-indexer
package types

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
