package args

import (
	"fmt"
	"slices"

	teetypes "github.com/masa-finance/tee-types/types"
)

// ValidateCapabilityForJobType validates that a capability is supported for the given job type
func ValidateCapabilityForJobType(jobType teetypes.JobType, capability teetypes.Capability) error {
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
