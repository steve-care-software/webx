package cardinalities

import (
	"github.com/steve-care-software/webx/engine/states/domain/hash"
)

// Builder represents the cardinality builder
type Builder interface {
	Create() Builder
	WithMin(min uint) Builder
	WithMax(max uint) Builder
	Now() (Cardinality, error)
}

// Cardinality represents the cardinality
type Cardinality interface {
	Hash() hash.Hash
	Min() uint
	HasMax() bool
	Max() bool
}
