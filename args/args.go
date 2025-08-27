package args

import (
	"encoding/json"
	"fmt"
	"strings"
)

// QueryTypeArgument provides a minimal structure to extract the QueryType (json "type")
// This is used across different job types to determine the specific capability being requested
type QueryTypeArgument struct {
	QueryType string `json:"type"`
}

// UnmarshalJSON implements custom JSON unmarshaling with normalization
func (q *QueryTypeArgument) UnmarshalJSON(data []byte) error {
	// Prevent infinite recursion
	type Alias QueryTypeArgument
	aux := &struct{ *Alias }{Alias: (*Alias)(q)}
	if err := json.Unmarshal(data, aux); err != nil {
		return fmt.Errorf("failed to unmarshal QueryType arguments: %w", err)
	}
	q.QueryType = strings.ToLower(aux.QueryType)
	return nil
}
