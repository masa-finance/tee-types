package linkedin

import (
	"github.com/masa-finance/tee-types/types/linkedin/experiences"
	"github.com/masa-finance/tee-types/types/linkedin/functions"
	"github.com/masa-finance/tee-types/types/linkedin/industries"
	"github.com/masa-finance/tee-types/types/linkedin/profile"
	"github.com/masa-finance/tee-types/types/linkedin/seniorities"
)

type linkedin struct {
	Profile     *profile.Profile
	Seniorities *seniorities.SenioritiesConfig
	Experiences *experiences.ExperiencesConfig
	Functions   *functions.FunctionsConfig
	Industries  *industries.IndustriesConfig
}

var LinkedIn = linkedin{
	Profile:     &profile.Profile{},
	Seniorities: &seniorities.Seniorities,
	Experiences: &experiences.Experiences,
	Functions:   &functions.Functions,
	Industries:  &industries.Industries,
}
