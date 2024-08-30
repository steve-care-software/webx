package validations

import "github.com/steve-care-software/webx/engine/domain/stacks/frames/variables"

type validation struct {
	variable variables.Variable
	isFail   bool
}

func createValidation(
	variable variables.Variable,
	isFail bool,
) Validation {
	out := validation{
		variable: variable,
		isFail:   isFail,
	}

	return &out
}

// Variable returns the variable
func (obj *validation) Variable() variables.Variable {
	return obj.variable
}

// IsFail returns true if fail, false otherwise
func (obj *validation) IsFail() bool {
	return obj.isFail
}
