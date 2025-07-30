package args

import (
	"fmt"
	"slices"

	teetypes "github.com/masa-finance/tee-types/types"
)

// jobCapabilityMap defines which capabilities are valid for each job type
var jobCapabilityMap = map[teetypes.JobType][]teetypes.Capability{
	// Twitter job types and their valid capabilities
	teetypes.TwitterJob: append(append(append(
		teetypes.TwitterCredentialCaps,
		teetypes.TwitterAPICaps...),
		teetypes.TwitterApifyCaps...),
		teetypes.CapSearchByFullArchive, // Elevated API capability
	),
	teetypes.TwitterCredentialJob: teetypes.TwitterCredentialCaps,
	teetypes.TwitterApiJob: append(
		teetypes.TwitterAPICaps,
		teetypes.CapSearchByFullArchive, // Elevated API capability
	),
	teetypes.TwitterApifyJob: teetypes.TwitterApifyCaps,

	// Web job capabilities
	teetypes.WebJob: teetypes.AlwaysAvailableWebCaps,

	// TikTok job capabilities
	teetypes.TiktokJob: teetypes.AlwaysAvailableTiktokCaps,

	// Telemetry job capabilities
	teetypes.TelemetryJob: teetypes.AlwaysAvailableTelemetryCaps,
}

// ValidateCapabilityForJobType validates that a capability is supported for the given job type
func ValidateCapabilityForJobType(jobType teetypes.JobType, capability teetypes.Capability) error {
	if capability == "" {
		// Empty capability is allowed for some job types
		return nil
	}

	validCaps, exists := jobCapabilityMap[jobType]
	if !exists {
		return fmt.Errorf("unknown job type: %s", jobType)
	}

	if !slices.Contains(validCaps, capability) {
		return fmt.Errorf("capability '%s' is not valid for job type '%s'. Valid capabilities: %v",
			capability, jobType, validCaps)
	}

	return nil
}

// GetValidCapabilitiesForJobType returns all valid capabilities for a given job type
func GetValidCapabilitiesForJobType(jobType teetypes.JobType) ([]teetypes.Capability, error) {
	validCaps, exists := jobCapabilityMap[jobType]
	if !exists {
		return nil, fmt.Errorf("unknown job type: %s", jobType)
	}

	// Return a copy to prevent external modification
	result := make([]teetypes.Capability, len(validCaps))
	copy(result, validCaps)
	return result, nil
}

// IsCapabilityValidForJobType checks if a capability is valid for a job type without returning an error
func IsCapabilityValidForJobType(jobType teetypes.JobType, capability teetypes.Capability) bool {
	return ValidateCapabilityForJobType(jobType, capability) == nil
}
