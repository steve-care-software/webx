package executions

import "github.com/steve-care-software/webx/engine/domain/grammars/blocks/lines/tokens/elements"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents an execution builder
type Builder interface {
	Create() Builder
	WithElements(elements elements.Elements) Builder
	WithFuncName(fnFlag string) Builder
	Now() (Execution, error)
}

// Execution represents an execution
type Execution interface {
	FuncName() string
	HasElements() bool
	Elements() elements.Elements
}
