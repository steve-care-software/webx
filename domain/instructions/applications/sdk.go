package applications

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithModule(module []byte) Builder
	WithName(name []byte) Builder
	Now() (Application, error)
}

// Application represents an application declaration
type Application interface {
	Module() []byte
	Name() []byte
}
