package args

import (
	"encoding/json"
	"fmt"
	"strings"

	teetypes "github.com/masa-finance/tee-types/types"
)

// TwitterSearchArguments defines args for Twitter searches
type TwitterSearchArguments struct {
	QueryType  string `json:"type"`  // Optional, type of search
	Query      string `json:"query"` // Username or search query
	Count      int    `json:"count"`
	StartTime  string `json:"start_time"`  // Optional ISO timestamp
	EndTime    string `json:"end_time"`    // Optional ISO timestamp
	MaxResults int    `json:"max_results"` // Optional, max number of results
	NextCursor string `json:"next_cursor"`
}

// UnmarshalJSON implements custom JSON unmarshaling with validation
func (t *TwitterSearchArguments) UnmarshalJSON(data []byte) error {
	type Alias TwitterSearchArguments
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(t),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return fmt.Errorf("failed to unmarshal Twitter arguments: %w", err)
	}

	// Normalize QueryType to lowercase
	t.QueryType = strings.ToLower(t.QueryType)

	return t.Validate()
}

// Validate validates the Twitter arguments (general validation)
func (t *TwitterSearchArguments) Validate() error {
	// note, query is not required for all capabilities

	if t.Count < 0 {
		return fmt.Errorf("count must be non-negative, got: %d", t.Count)
	}

	if t.MaxResults < 0 {
		return fmt.Errorf("max_results must be non-negative, got: %d", t.MaxResults)
	}

	return nil
}

// ValidateForJobType validates Twitter arguments for a specific job type
func (t *TwitterSearchArguments) ValidateForJobType(jobType teetypes.JobType) error {
	if err := t.Validate(); err != nil {
		return err
	}

	// Validate QueryType against job-specific capabilities
	return jobType.ValidateCapability(teetypes.Capability(t.QueryType))
}

// GetCapability returns the QueryType as a typed Capability
func (t *TwitterSearchArguments) GetCapability() teetypes.Capability {
	return teetypes.Capability(t.QueryType)
}

func (t *TwitterSearchArguments) IsSingleTweetOperation() bool {
	capability := t.GetCapability()
	return capability == teetypes.CapGetById
}

func (t *TwitterSearchArguments) IsMultipleTweetOperation() bool {
	capability := t.GetCapability()
	return capability == teetypes.CapSearchByQuery ||
		capability == teetypes.CapSearchByFullArchive ||
		capability == teetypes.CapGetHomeTweets ||
		capability == teetypes.CapGetForYouTweets ||
		capability == teetypes.CapGetTweets ||
		capability == teetypes.CapGetReplies ||
		capability == teetypes.CapGetMedia
}

func (t *TwitterSearchArguments) IsSingleProfileOperation() bool {
	capability := t.GetCapability()
	return capability == teetypes.CapGetProfileById ||
		capability == teetypes.CapSearchByProfile
}

func (t *TwitterSearchArguments) IsMultipleProfileOperation() bool {
	capability := t.GetCapability()
	return capability == teetypes.CapGetFollowing ||
		capability == teetypes.CapGetFollowers ||
		capability == teetypes.CapGetRetweeters
}

func (t *TwitterSearchArguments) IsSingleSpaceOperation() bool {
	capability := t.GetCapability()
	return capability == teetypes.CapGetSpace
}

func (t *TwitterSearchArguments) IsTrendsOperation() bool {
	capability := t.GetCapability()
	return capability == teetypes.CapGetTrends
}
