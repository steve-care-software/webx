package commands

import "github.com/steve-care-software/webx/domain/criterias"

type value struct {
	variable     criterias.Criteria
	constant     criterias.Criteria
	instructions criterias.Criteria
	execution    criterias.Criteria
}

func createValue(
	variable criterias.Criteria,
	constant criterias.Criteria,
	instructions criterias.Criteria,
	execution criterias.Criteria,
) Value {
	out := value{
		variable:     variable,
		constant:     constant,
		instructions: instructions,
		execution:    execution,
	}

	return &out
}

// Variable returns the variable
func (obj *value) Variable() criterias.Criteria {
	return obj.variable
}

// Constant returns the constant
func (obj *value) Constant() criterias.Criteria {
	return obj.constant
}

// Instructions returns the instructions
func (obj *value) Instructions() criterias.Criteria {
	return obj.instructions
}

// Execution returns the execution
func (obj *value) Execution() criterias.Criteria {
	return obj.execution
}
