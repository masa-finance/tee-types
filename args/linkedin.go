package args

import (
	"encoding/json"
	"fmt"
	"strings"

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
	if l.QueryType != "" {
		l.QueryType = strings.ToLower(l.QueryType)
	}

	return l.Validate()
}

// Validate validates the LinkedIn arguments
func (l *LinkedInArguments) Validate() error {
	if l.QueryType == "" {
		return fmt.Errorf("type is required")
	}

	// Validate query type
	validTypes := map[string]bool{
		"searchbyquery": true,
		"getprofile":    true,
	}
	if !validTypes[l.QueryType] {
		return fmt.Errorf("invalid type: %s, must be one of: searchbyquery, getprofile", l.QueryType)
	}

	if l.Query == "" {
		return fmt.Errorf("query is required")
	}

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
	return ValidateCapabilityForJobType(jobType, teetypes.Capability(l.QueryType))
}

// GetCapability returns the QueryType as a typed Capability
func (l *LinkedInArguments) GetCapability() teetypes.Capability {
	return teetypes.Capability(l.QueryType)
}

// LinkedInSearchArguments is an alias for LinkedInArguments for backward compatibility.
// Deprecated: use LinkedInArguments instead.
type LinkedInSearchArguments = LinkedInArguments
