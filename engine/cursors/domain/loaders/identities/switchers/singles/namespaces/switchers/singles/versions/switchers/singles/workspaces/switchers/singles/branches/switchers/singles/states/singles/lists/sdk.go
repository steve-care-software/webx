package lists

// Adapter represents a list adapter
type Adapter interface {
	ToBytes(ins List) ([]byte, error)
	ToInstance(data []byte) (List, error)
}

// Builder represents a list builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithResources(resources []string) Builder
	IsUnique() Builder
	Now() (List, error)
}

// List represents a list
type List interface {
	Name() string
	IsUnique() bool
	HasResources() bool
	Resources() []string
}
