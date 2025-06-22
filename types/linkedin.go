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