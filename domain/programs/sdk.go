package programs

import (
	"github.com/steve-care-software/syntax/domain/programs/instructions"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a program builder
type Builder interface {
	Create() Builder
	WithInstructions(instructions instructions.Instructions) Builder
	WithInputs(inputs []string) Builder
	WithOutputs(outputs []string) Builder
	Now() (Program, error)
}

// Program represents a program
type Program interface {
	Instructions() instructions.Instructions
	HasInputs() bool
	Inputs() []string
	HasOutputs() bool
	Outputs() []string
}
