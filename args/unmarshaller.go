package args

import (
	"encoding/json"
	"fmt"

	"github.com/masa-finance/tee-types/types"
)

// JobArgumentsInterface defines the interface that all job arguments must implement
type JobArgumentsInterface interface {
	Validate() error
	GetCapability() types.Capability
}

// TwitterJobArgumentsInterface extends JobArgumentsInterface for Twitter-specific methods
type TwitterJobArgumentsInterface interface {
	JobArgumentsInterface
	ValidateForJobType(jobType types.JobType) error
	IsNonTweetOperation() bool
}

// WebJobArgumentsInterface extends JobArgumentsInterface for Web-specific methods
type WebJobArgumentsInterface interface {
	JobArgumentsInterface
	IsDeepScrape() bool
	HasSelector() bool
	GetEffectiveMaxDepth() int
}

// TikTokJobArgumentsInterface extends JobArgumentsInterface for TikTok-specific methods
type TikTokJobArgumentsInterface interface {
	JobArgumentsInterface
	HasLanguagePreference() bool
	GetLanguageCode() string
}

// UnmarshalJobArguments unmarshals job arguments from a generic map into the appropriate typed struct
// This works with both tee-indexer and tee-worker JobArguments types
func UnmarshalJobArguments(jobType types.JobType, args map[string]any) (JobArgumentsInterface, error) {
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
func AsWebArguments(args JobArgumentsInterface) (WebJobArgumentsInterface, bool) {
	webArgs, ok := args.(*WebSearchArguments)
	if !ok {
		return nil, false
	}
	return webArgs, true
}

func AsTwitterArguments(args JobArgumentsInterface) (TwitterJobArgumentsInterface, bool) {
	twitterArgs, ok := args.(*TwitterSearchArguments)
	if !ok {
		return nil, false
	}
	return twitterArgs, true
}

func AsTikTokArguments(args JobArgumentsInterface) (TikTokJobArgumentsInterface, bool) {
	tiktokArgs, ok := args.(*TikTokTranscriptionArguments)
	if !ok {
		return nil, false
	}
	return tiktokArgs, true
}

func AsTelemetryArguments(args JobArgumentsInterface) (*TelemetryJobArguments, bool) {
	telemetryArgs, ok := args.(*TelemetryJobArguments)
	return telemetryArgs, ok
}
