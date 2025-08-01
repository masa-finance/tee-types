package args

import (
	"encoding/json"
	"fmt"

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
	GetLanguageCode() string
}

// LinkedInJobArguments extends JobArguments for LinkedIn-specific methods
type LinkedInJobArguments interface {
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

func unmarshalTikTokArguments(args map[string]any) (*TikTokTranscriptionArguments, error) {
	tiktokArgs := &TikTokTranscriptionArguments{}
	if err := unmarshalToStruct(args, tiktokArgs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal TikTok job arguments: %w", err)
	}
	return tiktokArgs, nil
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
		return nil, fmt.Errorf("Twitter job validation failed: %w", err)
	}

	return twitterArgs, nil
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
func AsWebArguments(args JobArguments) (WebJobArguments, bool) {
	webArgs, ok := args.(*WebSearchArguments)
	if !ok {
		return nil, false
	}
	return webArgs, true
}

func AsTwitterArguments(args JobArguments) (TwitterJobArguments, bool) {
	twitterArgs, ok := args.(*TwitterSearchArguments)
	if !ok {
		return nil, false
	}
	return twitterArgs, true
}

func AsTikTokArguments(args JobArguments) (TikTokJobArguments, bool) {
	tiktokArgs, ok := args.(*TikTokTranscriptionArguments)
	if !ok {
		return nil, false
	}
	return tiktokArgs, true
}

func AsTelemetryArguments(args JobArguments) (*TelemetryJobArguments, bool) {
	telemetryArgs, ok := args.(*TelemetryJobArguments)
	return telemetryArgs, ok
}

func AsLinkedInArguments(args JobArguments) (LinkedInJobArguments, bool) {
	linkedInArgs, ok := args.(*LinkedInArguments)
	if !ok {
		return nil, false
	}
	return linkedInArgs, true
}
