package frames

import (
	"github.com/steve-care-software/webx/engine/domain/stacks/frames/variables"
)

// NewFactory creates a new factory
func NewFactory() Factory {
	builder := NewBuilder()
	return createFactory(
		builder,
	)
}

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Factory represents a frame factory
type Factory interface {
	Create() (Frame, error)
}

// Builder represents a frame builder
type Builder interface {
	Create() Builder
	WithVariables(variables variables.Variables) Builder
	Now() (Frame, error)
}

// Frame represents a frame
type Frame interface {
	Fetch(name string) (variables.Variable, error)
	HasVariables() bool
	Variables() variables.Variables
}
