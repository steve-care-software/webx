package outputs

import "github.com/steve-care-software/syntax/domain/syntax/programs"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents an output builder
type Builder interface {
	Create() Builder
	WithProgram(program programs.Program) Builder
	WithScript(script []byte) Builder
	WithEngine(engine []byte) Builder
	Now() (Output, error)
}

// Output represents a compiled output
type Output interface {
	Program() programs.Program
	HasRemaining() bool
	Remaining() Remaining
}

// Remaining represents an output remaining
type Remaining interface {
	HasScript() bool
	Script() []byte
	HasEngine() bool
	Engine() []byte
}
