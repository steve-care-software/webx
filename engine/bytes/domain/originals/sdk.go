package originals

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents a namespace adapter
type Adapter interface {
	ToBytes(ins Original) ([]byte, error)
	ToInstance(data []byte) (Original, error)
}

// Builder represents an original builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithDescription(description string) Builder
	Now() (Original, error)
}

// Originals represents originals
type Originals interface {
	List() []Original
}

// Original represents an original namespace
type Original interface {
	Name() string
	Description() string
}
