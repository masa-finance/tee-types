package seniorities

import "github.com/masa-finance/tee-types/pkg/util"

// id represents a LinkedIn seniority level identifier
type Id string

// Seniority level constants
const (
	InTraining         Id = "100"
	EntryLevel         Id = "110"
	Senior             Id = "120"
	Strategic          Id = "130"
	EntryLevelManager  Id = "200"
	ExperiencedManager Id = "210"
	Director           Id = "220"
	VicePresident      Id = "300"
	CXO                Id = "310"
	Partner            Id = "320"
)

var All = util.NewSet(
	InTraining,
	EntryLevel,
	Senior,
	Strategic,
	EntryLevelManager,
	ExperiencedManager,
	Director,
	VicePresident,
	CXO,
	Partner,
)
