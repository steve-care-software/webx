package parameters

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
	WithToken(token string) ParameterBuilder
	WithIndex(index uint) ParameterBuilder
	WithName(name string) ParameterBuilder
	Now() (Parameter, error)
}

// Parameter represents an execution parameter
type Parameter interface {
	Token() string
	Index() uint
	Name() string
}
