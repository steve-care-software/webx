package compilers

import (
	"github.com/steve-care-software/webx/domain/criterias"
	"github.com/steve-care-software/webx/domain/grammars"
	"github.com/steve-care-software/webx/domain/instructions"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewElementsBuilder creates a new elements builder
func NewElementsBuilder() ElementsBuilder {
	return createElementsBuilder()
}

// NewElementBuilder creates a new element builder
func NewElementBuilder() ElementBuilder {
	return createElementBuilder()
}

// NewExecutionBuilder creates a new execution builder
func NewExecutionBuilder() ExecutionBuilder {
	return createExecutionBuilder()
}

// NewParametersBuilder creates a new parameters builder
func NewParametersBuilder() ParametersBuilder {
	return createParametersBuilder()
}

// NewParameterBuilder creates a new parameter builder
func NewParameterBuilder() ParameterBuilder {
	return createParameterBuilder()
}

// NewValueBuilder creates a new value builder
func NewValueBuilder() ValueBuilder {
	return createValueBuilder()
}

// Builder represents a compiler builder
type Builder interface {
	Create() Builder
	WithOutputs(outputs []string) Builder
	WithElements(elements Elements) Builder
	Now() (Compiler, error)
}

// Compiler represents a compiler
type Compiler interface {
	Elements() Elements
	HasOutputs() bool
	Outputs() []string
}

// ElementsBuilder represents an elements builder
type ElementsBuilder interface {
	Create() ElementsBuilder
	WithList(list []Element) ElementsBuilder
	Now() (Elements, error)
}

// Elements represents elements
type Elements interface {
	List() []Element
}

// ElementBuilder represents an element builder
type ElementBuilder interface {
	Create() ElementBuilder
	WithGrammar(grammar grammars.Grammar) ElementBuilder
	WithParameters(parameters Parameters) ElementBuilder
	WithExecution(execution Execution) ElementBuilder
	Now() (Element, error)
}

// Element represents an element
type Element interface {
	Grammar() grammars.Grammar
	Parameters() Parameters
	Execution() Execution
}

// ExecutionBuilder represents an execution builder
type ExecutionBuilder interface {
	Create() ExecutionBuilder
	WithParameter(parameter string) ExecutionBuilder
	WithInstructions(instructions instructions.Instructions) ExecutionBuilder
	Now() (Execution, error)
}

// Execution represents an instructions execution
type Execution interface {
	Parameter() string
	Instructions() instructions.Instructions
}

// ParametersBuilder represents a parameters builder
type ParametersBuilder interface {
	Create() ParametersBuilder
	WithList(list []Parameter) ParametersBuilder
	Now() (Parameters, error)
}

// Parameters represents parameters
type Parameters interface {
	List() []Parameter
}

// ParameterBuilder represents a parameter value
type ParameterBuilder interface {
	Create() ParameterBuilder
	WithName(name string) ParameterBuilder
	WithValue(value Value) ParameterBuilder
	Now() (Parameter, error)
}

// Parameter represents a parameter
type Parameter interface {
	Name() string
	Value() Value
}

// ValueBuilder represents a value builder
type ValueBuilder interface {
	Create() ValueBuilder
	WithConstant(constant string) ValueBuilder
	WithCriteria(criteria criterias.Criteria) ValueBuilder
	Now() (Value, error)
}

// Value represents a value
type Value interface {
	IsConstant() bool
	Constant() string
	IsCriteria() bool
	Criteria() criterias.Criteria
}
