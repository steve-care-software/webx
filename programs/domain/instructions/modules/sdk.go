package modules

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a module builder
type Builder interface {
	Create() Builder
	WithIndex(index uint) Builder
	WithName(name []byte) Builder
	Now() (Module, error)
}

// Module represents a module
type Module interface {
	Index() uint
	Name() []byte
}
