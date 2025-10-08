package profile

import (
	"time"

	"github.com/masa-finance/tee-types/pkg/util"
)

type ScraperMode string

const (
	ScraperModeShort     ScraperMode = "Short"
	ScraperModeFull      ScraperMode = "Full"
	ScraperModeFullEmail ScraperMode = "Full + email search"
)

var AllScraperModes = util.NewSet(ScraperModeShort, ScraperModeFull, ScraperModeFullEmail)

// Profile represents a complete profile response
type Profile struct {
	ID                      string            `json:"id"`
	PublicIdentifier        string            `json:"publicIdentifier"`
	URL                     string            `json:"linkedinUrl"`
	FirstName               string            `json:"firstName"`
	LastName                string            `json:"lastName"`
	Headline                string            `json:"headline"`
	About                   string            `json:"about"`
	OpenToWork              bool              `json:"openToWork"`
	Hiring                  bool              `json:"hiring"`
	Photo                   string            `json:"photo"`
	Premium                 bool              `json:"premium"`
	Influencer              bool              `json:"influencer"`
	Location                Location          `json:"location"`
	Verified                bool              `json:"verified"`
	RegisteredAt            time.Time         `json:"registeredAt"`
	TopSkills               string            `json:"topSkills"`
	ConnectionsCount        int               `json:"connectionsCount"`
	FollowerCount           int               `json:"followerCount"`
	CurrentPosition         []CurrentPosition `json:"currentPosition"`
	Experience              []Experience      `json:"experience"`
	Education               []Education       `json:"education"`
	Certifications          []Certification   `json:"certifications"`
	Projects                []Project         `json:"projects"`
	Volunteering            []Volunteering    `json:"volunteering"`
	ReceivedRecommendations []any             `json:"receivedRecommendations"` // we don't have examples of this data yet...
	Skills                  []Skill           `json:"skills"`
	Courses                 []Course          `json:"courses"`
	Publications            []Publication     `json:"publications"`
	Patents                 []any             `json:"patents"` // we don't have examples of this data yet...
	HonorsAndAwards         []HonorAndAward   `json:"honorsAndAwards"`
	Languages               []Language        `json:"languages"`
	Featured                any               `json:"featured"` // we don't have examples of this data yet...
	MoreProfiles            []MoreProfile     `json:"moreProfiles"`
}

// Location represents the location information
type Location struct {
	Text        string         `json:"Text"`
	CountryCode string         `json:"countryCode"`
	Parsed      ParsedLocation `json:"parsed"`
}

// ParsedLocation represents the parsed location details
type ParsedLocation struct {
	Text        string  `json:"text"`
	CountryCode string  `json:"countryCode"`
	RegionCode  *string `json:"regionCode"`
	Country     string  `json:"country"`
	CountryFull string  `json:"countryFull"`
	State       string  `json:"state"`
	City        string  `json:"city"`
}

// CurrentPosition represents current position information
type CurrentPosition struct {
	CompanyName string `json:"companyName"`
}

// Experience represents work experience
type Experience struct {
	Position             string    `json:"position"`
	Location             string    `json:"location"`
	EmploymentType       string    `json:"employmentType"`
	WorkplaceType        *string   `json:"workplaceType"`
	CompanyName          string    `json:"companyName"`
	CompanyURL           string    `json:"companyUrl"`
	CompanyID            string    `json:"companyId"`
	CompanyUniversalName string    `json:"companyUniversalName"`
	Duration             string    `json:"duration"`
	Description          string    `json:"description"`
	Skills               []string  `json:"skills"`
	StartDate            DateRange `json:"startDate"`
	EndDate              DateRange `json:"endDate"`
}

// DateRange represents a date range with month, year, and text
type DateRange struct {
	Month *string `json:"month,omitempty"`
	Year  *int    `json:"year,omitempty"`
	Text  string  `json:"text"`
}

// Education represents educational background
type Education struct {
	SchoolName   string    `json:"schoolName"`
	SchoolURL    string    `json:"schoolUrl"`
	Degree       string    `json:"degree"`
	FieldOfStudy *string   `json:"fieldOfStudy"`
	Skills       []string  `json:"skills"`
	StartDate    DateRange `json:"startDate"`
	EndDate      DateRange `json:"endDate"`
	Period       string    `json:"period"`
}

// Certification represents a certification
type Certification struct {
	Title        string `json:"title"`
	IssuedAt     string `json:"issuedAt"`
	IssuedBy     string `json:"issuedBy"`
	IssuedByLink string `json:"issuedByLink"`
}

// Project represents a project
type Project struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Duration    string    `json:"duration"`
	StartDate   DateRange `json:"startDate"`
	EndDate     DateRange `json:"endDate"`
}

// Volunteering represents volunteer experience
type Volunteering struct {
	Role             string     `json:"role"`
	Duration         string     `json:"duration"`
	StartDate        *DateRange `json:"startDate"`
	EndDate          *DateRange `json:"endDate"`
	OrganizationName string     `json:"organizationName"`
	OrganizationURL  *string    `json:"organizationUrl"`
	Cause            string     `json:"cause"`
}

// Skill represents a skill with optional positions and endorsements
type Skill struct {
	Name         string   `json:"name"`
	Positions    []string `json:"positions,omitempty"`
	Endorsements string   `json:"endorsements,omitempty"`
}

// Course represents a course
type Course struct {
	Title              string `json:"title"`
	AssociatedWith     string `json:"associatedWith"`
	AssociatedWithLink string `json:"associatedWithLink"`
}

// Publication represents a publication
type Publication struct {
	Title       string `json:"title"`
	PublishedAt string `json:"publishedAt"`
	Link        string `json:"link"`
}

// HonorAndAward represents an honor or award
type HonorAndAward struct {
	Title              string `json:"title"`
	IssuedBy           string `json:"issuedBy"`
	IssuedAt           string `json:"issuedAt"`
	Description        string `json:"description"`
	AssociatedWith     string `json:"associatedWith"`
	AssociatedWithLink string `json:"associatedWithLink"`
}

// Language represents a language with proficiency level
type Language struct {
	Name        string `json:"name"`
	Proficiency string `json:"proficiency"`
}

// MoreProfile represents a related profile
type MoreProfile struct {
	ID               string  `json:"id"`
	FirstName        string  `json:"firstName"`
	LastName         string  `json:"lastName"`
	Position         *string `json:"position,omitempty"`
	PublicIdentifier string  `json:"publicIdentifier"`
	URL              string  `json:"Url"`
}
