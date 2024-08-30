package lines

import (
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/executions"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/processors"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewLineBuilder creates a line builder
func NewLineBuilder() LineBuilder {
	return createLineBuilder()
}

// Builder represents a line builder
type Builder interface {
	Create() Builder
	WithList(list []Line) Builder
	Now() (Lines, error)
}

// Lines represents lines
type Lines interface {
	List() []Line
}

// LineBuilder represents a line builder
type LineBuilder interface {
	Create() LineBuilder
	WithTokens(tokens tokens.Tokens) LineBuilder
	WithProcessor(processor processors.Processor) LineBuilder
	WithSyscall(syscall executions.Execution) LineBuilder
	Now() (Line, error)
}

// Line represents a variable
type Line interface {
	Tokens() tokens.Tokens
	HasProcessor() bool
	Processor() processors.Processor
	HasSyscall() bool
	Syscall() executions.Execution
}
