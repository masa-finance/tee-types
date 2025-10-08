package linkedin

import (
	"github.com/masa-finance/tee-types/args/linkedin/profile"
)

type linkedin struct {
	Profile *profile.Arguments
}

var LinkedIn = linkedin{
	Profile: &profile.Arguments{},
}
