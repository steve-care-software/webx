package lines

import (
	"github.com/steve-care-software/webx/engine/domain/grammars/blocks/lines/executions"
	"github.com/steve-care-software/webx/engine/domain/grammars/blocks/lines/replacements"
)

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
	WithTokens(tokens []string) LineBuilder
	WithExecution(execution executions.Execution) LineBuilder
	WithReplacement(replacement replacements.Replacement) LineBuilder
	Now() (Line, error)
}

// Line represents a variable
type Line interface {
	Tokens() []string
	Execution() executions.Execution
	HasReplacement() bool
	Replacement() replacements.Replacement
}
