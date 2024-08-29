package frames

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/engine/domain/stacks/frames/variables"
)

type frame struct {
	variables variables.Variables
}

func createFrame() Frame {
	return createFrameInternally(nil)
}

func createFrameWithVariables(
	variables variables.Variables,
) Frame {
	return createFrameInternally(
		variables,
	)
}

func createFrameInternally(
	variables variables.Variables,
) Frame {
	out := frame{
		variables: variables,
	}

	return &out
}

// Fetch fetches a variable by name
func (obj *frame) Fetch(name string) (variables.Variable, error) {
	if !obj.HasVariables() {
		str := fmt.Sprintf("the frame contains no variable and therefore the variable (name: %s) could not be found", name)
		return nil, errors.New(str)
	}

	return obj.variables.Fetch(name)
}

// HasVariables returns true if there is variables, false otherwise
func (obj *frame) HasVariables() bool {
	return obj.variables != nil
}

// Variables returns the variables, if any
func (obj *frame) Variables() variables.Variables {
	return obj.variables
}
