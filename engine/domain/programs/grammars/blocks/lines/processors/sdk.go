package processors

import (
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/executions"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens/elements"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the processor builder
type Builder interface {
	Create() Builder
	WithExecution(execution executions.Execution) Builder
	WithReplacement(replacement elements.Element) Builder
	Now() (Processor, error)
}

// Processor represents a processor
type Processor interface {
	IsExecution() bool
	Execution() executions.Execution
	IsReplacement() bool
	Replacement() elements.Element
}
