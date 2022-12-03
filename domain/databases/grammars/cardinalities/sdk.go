package cardinalities

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	builder := NewBuilder()
	return createAdapter(builder)
}

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents a cardinality adapter
type Adapter interface {
	ToContent(ins Cardinality) ([]byte, error)
	ToCardinality(content []byte) (Cardinality, error)
}

// Builder represents a cardinality builder
type Builder interface {
	Create() Builder
	WithMin(min uint) Builder
	WithMax(max uint) Builder
	Now() (Cardinality, error)
}

// Cardinality represents cardinality
type Cardinality interface {
	Min() uint
	HasMax() bool
	Max() *uint
}
