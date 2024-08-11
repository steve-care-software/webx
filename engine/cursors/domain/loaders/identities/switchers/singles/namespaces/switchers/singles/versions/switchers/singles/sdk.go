package singles

// Adapter represents a single adapter
type Adapter interface {
	ToBytes(ins Single) ([]byte, error)
	ToInstance(data []byte) (Single, error)
}

// Builder represents a single builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithDescription(description string) Builder
	Now() (Single, error)
}

// Single represents a single namespace
type Single interface {
	Name() string
	Description() string
}
