package args

import (
	"fmt"
	"slices"

	teetypes "github.com/masa-finance/tee-types/types"
)

// ValidateCapabilityForJobType validates that a capability is supported for the given job type
func ValidateCapabilityForJobType(jobType teetypes.JobType, capability teetypes.Capability) error {
	if capability == "" {
		// Empty capability is allowed for some job types
		return nil
	}

	validCaps, exists := teetypes.JobCapabilityMap[jobType]
	if !exists {
		return fmt.Errorf("unknown job type: %s", jobType)
	}

	if !slices.Contains(validCaps, capability) {
		return fmt.Errorf("capability '%s' is not valid for job type '%s'. Valid capabilities: %v",
			capability, jobType, validCaps)
	}

	return nil
}

// IsCapabilityValidForJobType checks if a capability is valid for a job type without returning an error
func IsCapabilityValidForJobType(jobType teetypes.JobType, capability teetypes.Capability) bool {
	return ValidateCapabilityForJobType(jobType, capability) == nil
}
