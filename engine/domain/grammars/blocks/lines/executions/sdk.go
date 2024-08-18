package executions

import "github.com/steve-care-software/webx/engine/domain/grammars/tokens"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents an execution builder
type Builder interface {
	Create() Builder
	WithTokens(tokens tokens.Tokens) Builder
	WithFuncName(fnFlag string) Builder
	Now() (Execution, error)
}

// Execution represents an execution
type Execution interface {
	FuncName() string
	HasTokens() bool
	Tokens() tokens.Tokens
}
