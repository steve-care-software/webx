package parameters

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a parameter builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	IsInput() Builder
	Now() (Parameter, error)
}

// Parameter represents a parameter
type Parameter interface {
	Name() string
	IsInput() bool
}
