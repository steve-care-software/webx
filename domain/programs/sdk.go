package programs

import (
	program_instructions "github.com/steve-care-software/syntax/domain/programs/instructions"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a program builder
type Builder interface {
	Create() Builder
	WithInstructions(instructions program_instructions.Instructions) Builder
	WithInputs(inputs []string) Builder
	WithOutputs(outputs []string) Builder
	Now() (Program, error)
}

// Program represents a program
type Program interface {
	Instructions() program_instructions.Instructions
	HasInputs() bool
	Inputs() []string
	HasOutputs() bool
	Outputs() []string
}
