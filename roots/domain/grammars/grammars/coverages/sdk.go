package coverages

import (
	"github.com/steve-care-software/webx/roots/domain/grammars/grammars"
	"github.com/steve-care-software/webx/roots/domain/grammars/trees"
)

// NewBuilder initializes the builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewCoverageBuilder creates a new coverage builder
func NewCoverageBuilder() CoverageBuilder {
	return createCoverageBuilder()
}

// NewExecutionsBuilder creates a new executions builder
func NewExecutionsBuilder() ExecutionsBuilder {
	return createExecutionsBuilder()
}

// NewExecutionBuilder creates a new execution builder
func NewExecutionBuilder() ExecutionBuilder {
	return createExecutionBuilder()
}

// NewResultBuilder creates a new result builder
func NewResultBuilder() ResultBuilder {
	return createResultBuilder()
}

// Builder represents a coverages builder
type Builder interface {
	Create() Builder
	WithList(list []Coverage) Builder
	Now() (Coverages, error)
}

// Coverages represents coverages
type Coverages interface {
	List() []Coverage
	ContainsError() bool
}

// CoverageBuilder represents a coverage builder
type CoverageBuilder interface {
	Create() CoverageBuilder
	WithToken(token grammars.Token) CoverageBuilder
	WithExecutions(executions Executions) CoverageBuilder
	Now() (Coverage, error)
}

// Coverage represents a test coverage
type Coverage interface {
	Token() grammars.Token
	Executions() Executions
}

// ExecutionsBuilder represents an executions builder
type ExecutionsBuilder interface {
	Create() ExecutionsBuilder
	WithList(list []Execution) ExecutionsBuilder
	Now() (Executions, error)
}

// Executions represents executions
type Executions interface {
	List() []Execution
	ContainsError() bool
}

// ExecutionBuilder represents an execution builder
type ExecutionBuilder interface {
	Create() ExecutionBuilder
	WithExpectation(expectation grammars.Suite) ExecutionBuilder
	WithResult(result Result) ExecutionBuilder
	Now() (Execution, error)
}

// Execution represents a suite's execution
type Execution interface {
	Expectation() grammars.Suite
	Result() Result
}

// ResultBuilder represents a result builder
type ResultBuilder interface {
	Create() ResultBuilder
	WithTree(tree trees.Tree) ResultBuilder
	WithError(error string) ResultBuilder
	Now() (Result, error)
}

// Result represents an expectation's result
type Result interface {
	IsTree() bool
	Tree() trees.Tree
	IsError() bool
	Error() string
}
