package outputs

import "github.com/steve-care-software/syntax/domain/syntax/programs"

// NewBuilder represents an output builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents an output builder
type Builder interface {
	Create() Builder
	WithProgram(program programs.Program) Builder
	WithRemaining(remaining []byte) Builder
	Now() (Output, error)
}

// Output represents an output
type Output interface {
	Program() programs.Program
	HasRemaining() bool
	Remaining() []byte
}
