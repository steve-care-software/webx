package resources

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/queries/conditions/pointers"
)

// NewBuilder creates a new resource builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents a resource builder
type Builder interface {
	Create() Builder
	WithField(field pointers.Pointer) Builder
	WithValue(value interface{}) Builder
	Now() (Resource, error)
}

// Resource represents a resource
type Resource interface {
	Hash() hash.Hash
	IsField() bool
	Field() pointers.Pointer
	IsValue() bool
	Value() interface{}
}
