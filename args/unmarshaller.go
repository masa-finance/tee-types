package args

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/masa-finance/tee-types/types"
)

// JobArguments defines the interface that all job arguments must implement
type JobArguments interface {
	Validate() error
	GetCapability() types.Capability
}

// TwitterJobArguments extends JobArguments for Twitter-specific methods
type TwitterJobArguments interface {
	JobArguments
	ValidateForJobType(jobType types.JobType) error
	IsSingleTweetOperation() bool
	IsMultipleTweetOperation() bool
	IsSingleProfileOperation() bool
	IsMultipleProfileOperation() bool
	IsSingleSpaceOperation() bool
	IsTrendsOperation() bool
}

// WebJobArguments extends JobArguments for Web-specific methods
type WebJobArguments interface {
	JobArguments
	ValidateForJobType(jobType types.JobType) error
	IsDeepScrape() bool
	HasSelector() bool
	GetEffectiveMaxDepth() int
}

// TikTokJobArguments extends JobArguments for TikTok-specific methods
type TikTokJobArguments interface {
	JobArguments
	ValidateForJobType(jobType types.JobType) error
	HasLanguagePreference() bool
	GetVideoURL() string
	GetLanguageCode() string
}

// LinkedInJobArguments extends JobArguments for LinkedIn-specific methods
type LinkedInJobArguments interface {
	JobArguments
	ValidateForJobType(jobType types.JobType) error
}

// RedditJobArguments extends JobArguments for Reddit-specific methods
type RedditJobArguments interface {
	JobArguments
	ValidateForJobType(jobType types.JobType) error
}

// UnmarshalJobArguments unmarshals job arguments from a generic map into the appropriate typed struct
// This works with both tee-indexer and tee-worker JobArguments types
func UnmarshalJobArguments(jobType types.JobType, args map[string]any) (JobArguments, error) {
	switch jobType {
	case types.WebJob:
		return unmarshalWebArguments(args)

	case types.TiktokJob:
		return unmarshalTikTokArguments(args)

	case types.TwitterJob, types.TwitterCredentialJob, types.TwitterApiJob, types.TwitterApifyJob:
		return unmarshalTwitterArguments(jobType, args)

	case types.LinkedInJob:
		return unmarshalLinkedInArguments(jobType, args)

	case types.RedditJob:
		return unmarshalRedditArguments(jobType, args)

	case types.TelemetryJob:
		return &TelemetryJobArguments{}, nil

	default:
		return nil, fmt.Errorf("unknown job type: %s", jobType)
	}
}

// Helper functions for unmarshaling specific argument types
func unmarshalWebArguments(args map[string]any) (*WebSearchArguments, error) {
	webArgs := &WebSearchArguments{}
	if err := unmarshalToStruct(args, webArgs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal web job arguments: %w", err)
	}
	return webArgs, nil
}

func unmarshalTikTokArguments(args map[string]any) (JobArguments, error) {
	// Unmarshal minimally to read QueryType like we do for Twitter
	minimal := &TikTokArguments{}
	if err := unmarshalToStruct(args, minimal); err != nil {
		return nil, fmt.Errorf("failed to unmarshal TikTok arguments: %w", err)
	}
	capability := types.Capability(strings.ToLower(minimal.QueryType))
	if capability == types.Capability("") {
		defaultCap, exists := types.JobDefaultCapabilityMap[types.TiktokJob]
		if !exists {
			return nil, fmt.Errorf("no default capability configured for job type: %s", types.TiktokJob)
		}
		capability = defaultCap
	}

	switch capability {
	case types.CapSearchByQuery:
		searchArgs := &TikTokSearchByQueryArguments{}
		if err := unmarshalToStruct(args, searchArgs); err != nil {
			return nil, fmt.Errorf("failed to unmarshal TikTok searchbyquery arguments: %w", err)
		}
		if err := searchArgs.ValidateForJobType(types.TiktokJob); err != nil {
			return nil, fmt.Errorf("tiktok job validation failed: %w", err)
		}
		return searchArgs, nil
	case types.CapSearchByTrending:
		searchArgs := &TikTokSearchByTrendingArguments{}
		if err := unmarshalToStruct(args, searchArgs); err != nil {
			return nil, fmt.Errorf("failed to unmarshal TikTok searchbytrending arguments: %w", err)
		}
		if err := searchArgs.ValidateForJobType(types.TiktokJob); err != nil {
			return nil, fmt.Errorf("tiktok job validation failed: %w", err)
		}
		return searchArgs, nil
	case types.CapTranscription:
		transcriptionArgs := &TikTokTranscriptionArguments{}
		if err := unmarshalToStruct(args, transcriptionArgs); err != nil {
			return nil, fmt.Errorf("failed to unmarshal TikTok transcription arguments: %w", err)
		}
		if err := transcriptionArgs.ValidateForJobType(types.TiktokJob); err != nil {
			return nil, fmt.Errorf("tiktok job validation failed: %w", err)
		}
		return transcriptionArgs, nil
	default:
		return nil, fmt.Errorf("unknown tiktok type: %s", capability)
	}
}

func unmarshalTwitterArguments(jobType types.JobType, args map[string]any) (*TwitterSearchArguments, error) {
	twitterArgs := &TwitterSearchArguments{}
	if err := unmarshalToStruct(args, twitterArgs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal Twitter job arguments: %w", err)
	}

	// If no QueryType is specified, use the default capability for this job type
	if twitterArgs.QueryType == "" {
		if defaultCap, exists := types.JobDefaultCapabilityMap[jobType]; exists {
			twitterArgs.QueryType = string(defaultCap)
		}
	}

	// Perform job-type-specific validation for Twitter
	if err := twitterArgs.ValidateForJobType(jobType); err != nil {
		return nil, fmt.Errorf("twitter job validation failed: %w", err)
	}

	return twitterArgs, nil
}

func unmarshalLinkedInArguments(jobType types.JobType, args map[string]any) (*LinkedInArguments, error) {
	linkedInArgs := &LinkedInArguments{}
	if err := unmarshalToStruct(args, linkedInArgs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal LinkedIn job arguments: %w", err)
	}

	// If no QueryType is specified, use the default capability for this job type
	if linkedInArgs.QueryType == "" {
		if defaultCap, exists := types.JobDefaultCapabilityMap[jobType]; exists {
			linkedInArgs.QueryType = string(defaultCap)
		}
	}

	// Perform job-type-specific validation for LinkedIn
	if err := linkedInArgs.ValidateForJobType(jobType); err != nil {
		return nil, fmt.Errorf("linkedin job validation failed: %w", err)
	}

	return linkedInArgs, nil
}

func unmarshalRedditArguments(jobType types.JobType, args map[string]any) (*RedditArguments, error) {
	redditArgs := &RedditArguments{}
	if err := unmarshalToStruct(args, redditArgs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal Reddit job arguments: %w", err)
	}

	// If no QueryType is specified, use the default capability for this job type
	if redditArgs.QueryType == "" {
		if defaultCap, exists := types.JobDefaultCapabilityMap[jobType]; exists {
			redditArgs.QueryType = types.RedditQueryType(defaultCap)
		}
	}

	// Perform job-type-specific validation for Reddit
	if err := redditArgs.ValidateForJobType(jobType); err != nil {
		return nil, fmt.Errorf("reddit job validation failed: %w", err)
	}

	return redditArgs, nil
}

// unmarshalToStruct converts a map[string]any to a struct using JSON marshal/unmarshal
// This provides the same functionality as the existing JobArguments.Unmarshal methods
func unmarshalToStruct(args map[string]any, target any) error {
	// Use JSON marshal/unmarshal for conversion - this triggers our custom UnmarshalJSON methods
	data, err := json.Marshal(args)
	if err != nil {
		return fmt.Errorf("failed to marshal arguments: %w", err)
	}

	if err := json.Unmarshal(data, target); err != nil {
		return fmt.Errorf("failed to unmarshal arguments: %w", err)
	}

	return nil
}

// TelemetryJobArguments for telemetry jobs (simple case)
type TelemetryJobArguments struct{}

func (t *TelemetryJobArguments) Validate() error {
	return nil
}

func (t *TelemetryJobArguments) GetCapability() types.Capability {
	return types.CapTelemetry
}

// Type assertion helpers
func AsWebArguments(args JobArguments) (*WebSearchArguments, bool) {
	webArgs, ok := args.(*WebSearchArguments)
	return webArgs, ok
}

func AsTwitterArguments(args JobArguments) (TwitterJobArguments, bool) {
	twitterArgs, ok := args.(*TwitterSearchArguments)
	if !ok {
		return nil, false
	}
	return twitterArgs, true
}

// Use specific helpers for TikTok argument types:
// - AsTikTokTranscriptionArguments
// - AsTikTokSearchByQueryArguments
// - AsTikTokSearchByTrendingArguments

func AsTikTokTranscriptionArguments(args JobArguments) (*TikTokTranscriptionArguments, bool) {
	v, ok := args.(*TikTokTranscriptionArguments)
	return v, ok
}

func AsTikTokSearchByQueryArguments(args JobArguments) (*TikTokSearchByQueryArguments, bool) {
	v, ok := args.(*TikTokSearchByQueryArguments)
	return v, ok
}

func AsTikTokSearchByTrendingArguments(args JobArguments) (*TikTokSearchByTrendingArguments, bool) {
	v, ok := args.(*TikTokSearchByTrendingArguments)
	return v, ok
}
