package lists

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/lists/deletes"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/lists/inserts"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the list adapter
type Adapter interface {
	ToBytes(ins List) ([]byte, error)
	ToInstance(bytes []byte) (List, error)
}

// Builder represents a list builder
type Builder interface {
	Create() Builder
	WithInsert(insert inserts.Insert) Builder
	WithDelete(delete deletes.Delete) Builder
	Now() (List, error)
}

// List represents a list assignable
type List interface {
	Hash() hash.Hash
	IsInsert() bool
	Insert() inserts.Insert
	IsDelete() bool
	Delete() deletes.Delete
}
