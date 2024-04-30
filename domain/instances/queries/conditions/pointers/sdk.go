package pointers

import (
	"github.com/steve-care-software/datastencil/domain/hash"
)

// NewBuilder creates a new pointer builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents a pointer builder
type Builder interface {
	Create() Builder
	WithEntity(entity string) Builder
	WithField(field string) Builder
	Now() (Pointer, error)
}

// Pointer represents a field pointer
type Pointer interface {
	Hash() hash.Hash
	Entity() string
	Field() string
}
