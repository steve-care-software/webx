package stackframes

import (
	"github.com/steve-care-software/webx/engine/domain/stacks"
	"github.com/steve-care-software/webx/engine/domain/stacks/frames"
	"github.com/steve-care-software/webx/engine/domain/stacks/frames/variables"
)

// NewFactory creates a new factory
func NewFactory() Factory {
	stackFactory := stacks.NewFactory()
	stackBuilder := stacks.NewBuilder()
	frameBuilder := frames.NewBuilder()
	variablesBuilder := variables.NewBuilder()
	variableBuilder := variables.NewVariableBuilder()
	return cretaeFactory(
		stackFactory,
		stackBuilder,
		frameBuilder,
		variablesBuilder,
		variableBuilder,
	)
}

// Factory represents the application factory
type Factory interface {
	Create() (Application, error)
}

// Application represents a stackframe application
type Application interface {
	Root()                                    // selects the root stack
	Navigate(path []string) error             // navigate to the stack where the path points
	Fetch() (stacks.Stack, error)             // fetches the stack where our cursor points to
	Push(variables variables.Variables) error // push the stack where our cursor points to
	Pop() error                               // pop the stack where our cursor points to
}
