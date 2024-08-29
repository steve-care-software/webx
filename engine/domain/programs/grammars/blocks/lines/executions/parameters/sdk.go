package parameters

import (
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/executions/parameters/values"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewParameterBuilder creates a new parameter builder
func NewParameterBuilder() ParameterBuilder {
	return createParameterBuilder()
}

// Builder represents the builder
type Builder interface {
	Create() Builder
	WithList(list []Parameter) Builder
	Now() (Parameters, error)
}

// Parameters represents parameters
type Parameters interface {
	List() []Parameter
}

// ParameterBuilder represents the parameter builder
type ParameterBuilder interface {
	Create() ParameterBuilder
	WithName(name string) ParameterBuilder
	WithValue(value values.Value) ParameterBuilder
	Now() (Parameter, error)
}

// Parameter represents an execution parameter
type Parameter interface {
	Name() string
	Value() values.Value
}
