package cardinalities

import "github.com/steve-care-software/webx/domain/databases/entities"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a cardinality builder
type Builder interface {
	Create() Builder
	WithEntity(entity entities.Entity) Builder
	WithMin(min uint) Builder
	WithMax(max uint) Builder
	Now() (Cardinality, error)
}

// Cardinality represents cardinality
type Cardinality interface {
	Entity() entities.Entity
	Min() uint
	HasMax() bool
	Max() *uint
}
