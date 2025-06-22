package args

// LinkedInSearchArguments defines args for LinkedIn searches
type LinkedInSearchArguments struct {
	QueryType      string   `json:"type"`                       // "searchbyquery", "getprofile"
	Query          string   `json:"query"`                      // Keywords for search or username for profile
	NetworkFilters []string `json:"network_filters,omitempty"` // ["F", "S", "O"] - First, Second, Other (default: all)
	MaxResults     int      `json:"max_results"`                // Maximum number of results to return
	Start          int      `json:"start"`                      // Pagination start offset
}