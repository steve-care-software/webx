package coverages

import "github.com/steve-care-software/syntax/domain/syntax/bytes/grammars"

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

// NewLineBuilder creates a new line builder
func NewLineBuilder() LineBuilder {
	return createLineBuilder()
}

// NewElementBuilder creates a new element builder
func NewElementBuilder() ElementBuilder {
	return createElementBuilder()
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
}

// ExecutionBuilder represents an execution builder
type ExecutionBuilder interface {
	Create() ExecutionBuilder
	WithSuite(suite grammars.Suite) ExecutionBuilder
	WithLine(line Line) ExecutionBuilder
	Now() (Execution, error)
}

// Execution represents a suite's execution
type Execution interface {
	Suite() grammars.Suite
	Line() Line
}

// LineBuilder represents a line builder
type LineBuilder interface {
	Create() LineBuilder
	WithList(list []Element) LineBuilder
	Now() (Line, error)
}

// Line represents a coverage line
type Line interface {
	List() []Element
}

// ElementBuilder represents an element builder
type ElementBuilder interface {
	Create() ElementBuilder
	WithName(name string) ElementBuilder
	WithValue(value []byte) ElementBuilder
	Now() (Element, error)
}

// Element represents a coverage element
type Element interface {
	Name() string
	HasValue() bool
	Value() []byte
}
