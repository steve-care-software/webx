package compilers

import "github.com/steve-care-software/syntax/domain/syntax/criterias"

type replacement struct {
	name     []byte
	criteria criterias.Criteria
}

func createReplacement(
	name []byte,
	criteria criterias.Criteria,
) Replacement {
	out := replacement{
		name:     name,
		criteria: criteria,
	}

	return &out
}

// Name returns the name
func (obj *replacement) Name() []byte {
	return obj.name
}

// Criteria returns the criteria
func (obj *replacement) Criteria() criterias.Criteria {
	return obj.criteria
}
