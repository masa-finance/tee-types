package profile

import (
	"encoding/json"
	"errors"
	"fmt"

	teetypes "github.com/masa-finance/tee-types/types"
	"github.com/masa-finance/tee-types/types/linkedin/experiences"
	"github.com/masa-finance/tee-types/types/linkedin/functions"
	"github.com/masa-finance/tee-types/types/linkedin/industries"
	"github.com/masa-finance/tee-types/types/linkedin/profile"
	"github.com/masa-finance/tee-types/types/linkedin/seniorities"
)

var (
	ErrScraperModeNotSupported       = errors.New("scraper mode not supported")
	ErrMaxItemsTooLarge              = errors.New("max items must be less than or equal to 100")
	ErrYearsOfExperienceNotSupported = errors.New("years of experience not supported")
	ErrSeniorityLevelNotSupported    = errors.New("seniority level not supported")
	ErrFunctionNotSupported          = errors.New("function not supported")
	ErrIndustryNotSupported          = errors.New("industry not supported")
)

const (
	DefaultMaxItems    = 10
	DefaultScraperMode = profile.ScraperModeShort
	MaxItems           = 100
)

// Arguments defines args for LinkedIn profile operations
type Arguments struct {
	QueryType             teetypes.Capability `json:"type"`
	ScraperMode           profile.ScraperMode `json:"scraperMode"`
	Query                 string              `json:"searchQuery"`
	MaxItems              uint                `json:"maxItems"`
	Locations             []string            `json:"locations"`
	CurrentComanies       []string            `json:"currentCompanies"`
	PastCompanies         []string            `json:"pastCompanies"`
	CurrentJobTitles      []string            `json:"currentJobTitles"`
	PastJobTitles         []string            `json:"pastJobTitles"`
	Schools               []string            `json:"schools"`
	YearsOfExperience     []experiences.Id    `json:"yearsOfExperienceIds"`
	YearsAtCurrentCompany []experiences.Id    `json:"yearsAtCurrentCompanyIds"`
	SeniorityLevels       []seniorities.Id    `json:"seniorityLevelIds"`
	Functions             []functions.Id      `json:"functionIds"`
	Industries            []industries.Id     `json:"industryIds"`
	FirstNames            []string            `json:"firstNames"`
	LastNames             []string            `json:"lastNames"`
	RecentlyChangedJobs   bool                `json:"recentlyChangedJobs"`
	StartPage             uint                `json:"startPage"`
}

func (a *Arguments) UnmarshalJSON(data []byte) error {
	type Alias Arguments
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(a),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return fmt.Errorf("failed to unmarshal LinkedIn profile arguments: %w", err)
	}

	a.setDefaultValues()

	return a.Validate()
}

func (a *Arguments) setDefaultValues() {
	if a.MaxItems == 0 {
		a.MaxItems = DefaultMaxItems
	}
	if a.ScraperMode == "" {
		a.ScraperMode = DefaultScraperMode
	}
}

func (a *Arguments) Validate() error {
	var errs []error

	if a.MaxItems > MaxItems {
		errs = append(errs, ErrMaxItemsTooLarge)
	}
	if !profile.AllScraperModes.Contains(a.ScraperMode) {
		errs = append(errs, ErrScraperModeNotSupported)
	}
	for _, yoe := range a.YearsOfExperience {
		if !experiences.All.Contains(yoe) {
			errs = append(errs, fmt.Errorf("%w: %v", ErrYearsOfExperienceNotSupported, yoe))
		}
	}
	for _, yac := range a.YearsAtCurrentCompany {
		if !experiences.All.Contains(yac) {
			errs = append(errs, fmt.Errorf("%w: %v", ErrYearsOfExperienceNotSupported, yac))
		}
	}
	for _, sl := range a.SeniorityLevels {
		if !seniorities.All.Contains(sl) {
			errs = append(errs, fmt.Errorf("%w: %v", ErrSeniorityLevelNotSupported, sl))
		}
	}
	for _, f := range a.Functions {
		if !functions.All.Contains(f) {
			errs = append(errs, fmt.Errorf("%w: %v", ErrFunctionNotSupported, f))
		}
	}
	for _, i := range a.Industries {
		if !industries.All.Contains(i) {
			errs = append(errs, fmt.Errorf("%w: %v", ErrIndustryNotSupported, i))
		}
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}

func (a *Arguments) GetCapability() teetypes.Capability {
	return a.QueryType
}

func (a *Arguments) ValidateForJobType(jobType teetypes.JobType) error {
	if err := a.Validate(); err != nil {
		return err
	}

	return jobType.ValidateCapability(a.QueryType)
}
