package args

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/masa-finance/tee-types/pkg/util"
	teetypes "github.com/masa-finance/tee-types/types"
)

// LinkedInArguments defines args for LinkedIn operations
type LinkedInArguments struct {
	QueryType        string   `json:"type"`  // "searchbyquery", "getprofile"
	Query            string   `json:"query"` // Keywords for search or username for profile
	PublicIdentifier string   `json:"public_identifier,omitempty"`
	NetworkFilters   []string `json:"network_filters,omitempty"` // ["F", "S", "O"] - First, Second, Other (default: all)
	MaxResults       int      `json:"max_results"`               // Maximum number of results to return
	Start            int      `json:"start"`                     // Pagination start offset
}

// UnmarshalJSON implements custom JSON unmarshaling with validation
func (l *LinkedInArguments) UnmarshalJSON(data []byte) error {
	// Prevent infinite recursion (you call json.Unmarshal which then calls `UnmarshalJSON`, which then calls `json.Unmarshal`...)
	type Alias LinkedInArguments
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(l),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return fmt.Errorf("failed to unmarshal LinkedIn arguments: %w", err)
	}

	// Normalize QueryType to lowercase
	l.QueryType = strings.ToLower(l.QueryType)

	return l.Validate()
}

// Validate validates the LinkedIn arguments (general validation)
func (l *LinkedInArguments) Validate() error {
	// Note: QueryType is not required for all capabilities, similar to Twitter pattern
	// Query is also not required for all capabilities

	if l.MaxResults < 0 {
		return fmt.Errorf("max_results must be non-negative, got: %d", l.MaxResults)
	}

	if l.Start < 0 {
		return fmt.Errorf("start must be non-negative, got: %d", l.Start)
	}

	return nil
}

// ValidateForJobType validates LinkedIn arguments for a specific job type
func (l *LinkedInArguments) ValidateForJobType(jobType teetypes.JobType) error {
	if err := l.Validate(); err != nil {
		return err
	}

	// Validate QueryType against job-specific capabilities
	return jobType.ValidateCapability(teetypes.Capability(l.QueryType))
}

// GetCapability returns the QueryType as a typed Capability
func (l *LinkedInArguments) GetCapability() teetypes.Capability {
	return teetypes.Capability(l.QueryType)
}

// IsSearchOperation returns true if this is a search operation
func (l *LinkedInArguments) IsSearchOperation() bool {
	capability := l.GetCapability()
	return capability == teetypes.CapSearchByQuery
}

// IsProfileOperation returns true if this is a profile operation
func (l *LinkedInArguments) IsProfileOperation() bool {
	capability := l.GetCapability()
	return capability == teetypes.CapGetProfile
}

// HasNetworkFilters returns true if network filters are specified
func (l *LinkedInArguments) HasNetworkFilters() bool {
	return len(l.NetworkFilters) > 0
}

// GetEffectiveMaxResults returns the effective maximum results, defaulting to a reasonable limit
func (l *LinkedInArguments) GetEffectiveMaxResults() int {
	return util.Max(l.MaxResults, 10)
}

// LinkedInSearchArguments is an alias for LinkedInArguments for backward compatibility.
// Deprecated: use LinkedInArguments instead.
type LinkedInSearchArguments = LinkedInArguments
