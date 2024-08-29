package frames

import (
	"github.com/steve-care-software/webx/engine/domain/hash"
	"github.com/steve-care-software/webx/engine/domain/stacks/frames/variables"
)

// Builder represents a frame builder
type Builder interface {
	Create() Builder
	WithVariables(variables variables.Variables) Builder
	Now() (Frame, error)
}

// Frame represents a frame
type Frame interface {
	Hash() hash.Hash
	Fetch(name string) (variables.Variable, error)
	HasVariables() bool
	Variables() (variables.Variables, error)
}
