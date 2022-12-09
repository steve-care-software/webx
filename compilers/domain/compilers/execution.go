package compilers

import "github.com/steve-care-software/webx/programs/domain/instructions"

type execution struct {
	parameter    string
	instructions instructions.Instructions
}

func createExecution(
	parameter string,
	instructions instructions.Instructions,
) Execution {
	out := execution{
		parameter:    parameter,
		instructions: instructions,
	}

	return &out
}

// Parameter returns the parameter
func (obj *execution) Parameter() string {
	return obj.parameter
}

// Instructions returns the instructions
func (obj *execution) Instructions() instructions.Instructions {
	return obj.instructions
}
