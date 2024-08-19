package lines

import (
	"github.com/steve-care-software/webx/engine/domain/grammars/blocks/lines/executions"
	"github.com/steve-care-software/webx/engine/domain/grammars/blocks/lines/tokens"
	"github.com/steve-care-software/webx/engine/domain/grammars/blocks/lines/tokens/elements"
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
	WithExecution(execution executions.Execution) LineBuilder
	WithReplacement(replacement elements.Element) LineBuilder
	Now() (Line, error)
}

// Line represents a variable
type Line interface {
	Tokens() tokens.Tokens
	HasExecution() bool
	Execution() executions.Execution
	HasReplacement() bool
	Replacement() elements.Element
}
