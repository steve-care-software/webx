package parameters

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the parameter builder
type Builder interface {
	Create() Builder
	WithElement(element string) Builder
	WithIndex(index uint) Builder
	WithName(name string) Builder
	Now() (Parameter, error)
}

// Parameter represents an execution parameter
type Parameter interface {
	Element() string
	Index() uint
	Name() string
}
