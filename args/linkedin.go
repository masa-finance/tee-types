package args

// LinkedInArguments defines args for LinkedIn operations
type LinkedInArguments struct {
	QueryType        string   `json:"type"`  // "searchbyquery", "getprofile"
	Query            string   `json:"query"` // Keywords for search or username for profile
	PublicIdentifier string   `json:"public_identifier,omitempty"`
	NetworkFilters   []string `json:"network_filters,omitempty"` // ["F", "S", "O"] - First, Second, Other (default: all)
	MaxResults       int      `json:"max_results"`               // Maximum number of results to return
	Start            int      `json:"start"`                     // Pagination start offset
}

// LinkedInSearchArguments is an alias for LinkedInArguments for backward compatibility.
// Deprecated: use LinkedInArguments instead.
type LinkedInSearchArguments = LinkedInArguments
