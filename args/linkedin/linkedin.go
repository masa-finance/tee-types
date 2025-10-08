package linkedin

import (
	"github.com/masa-finance/tee-types/args/linkedin/profile"
)

// TODO: clean this up, you can't surface types in a struct...
type linkedin struct {
	Profile *profile.Arguments
}

var LinkedIn = linkedin{
	Profile: &profile.Arguments{},
}
