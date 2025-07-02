// Package types provides shared types between tee-worker and tee-indexer
package types

// LinkedInProfileResult defines the structure of a LinkedIn profile search result
type LinkedInProfileResult struct {
	PublicIdentifier string `json:"public_identifier"` // Username/slug in profile URL
	URN              string `json:"urn"`               // LinkedIn's unique resource name
	FullName         string `json:"full_name"`         // Person's full name
	Headline         string `json:"headline"`          // Professional headline/title
	Location         string `json:"location"`          // Geographic location
	ProfileURL       string `json:"profile_url"`       // Full LinkedIn profile URL
	Degree           string `json:"degree,omitempty"`  // Connection degree (1st, 2nd, etc.)
}

// Experience defines the structure for a single entry in a user's work experience
type Experience struct {
	Title       string `json:"title"`
	CompanyName string `json:"company_name"`
	Location    string `json:"location,omitempty"`
	StartDate   string `json:"start_date,omitempty"`
	EndDate     string `json:"end_date,omitempty"`
	Description string `json:"description,omitempty"`
}

// Education defines the structure for a single entry in a user's education history
type Education struct {
	SchoolName   string `json:"school_name"`
	DegreeName   string `json:"degree_name,omitempty"`
	FieldOfStudy string `json:"field_of_study,omitempty"`
	StartDate    string `json:"start_date,omitempty"`
	EndDate      string `json:"end_date,omitempty"`
	Description  string `json:"description,omitempty"`
}

// Skill defines the structure for a single skill entry
type Skill struct {
	Name string `json:"name"`
}

// LinkedInFullProfileResult defines the structure for a detailed LinkedIn profile
type LinkedInFullProfileResult struct {
	PublicIdentifier  string       `json:"public_identifier"`
	URN               string       `json:"urn"`
	FullName          string       `json:"full_name"`
	Headline          string       `json:"headline"`
	Location          string       `json:"location"`
	Summary           string       `json:"summary,omitempty"`
	ProfilePictureURL string       `json:"profile_picture_url,omitempty"`
	Experiences       []Experience `json:"experiences,omitempty"`
	Education         []Education  `json:"education,omitempty"`
	Skills            []Skill      `json:"skills,omitempty"`
}
