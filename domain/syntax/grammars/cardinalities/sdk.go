package cardinalities

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a cardinality builder
type Builder interface {
	Create() Builder
	WithMin(min uint) Builder
	WithMax(max uint) Builder
	Now() (Cardinality, error)
}

// Cardinality represents a cardinality
type Cardinality interface {
	Min() uint
	HasMax() bool
	Max() *uint
}
