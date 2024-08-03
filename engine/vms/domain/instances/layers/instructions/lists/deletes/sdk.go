package deletes

import "github.com/steve-care-software/webx/engine/hashes/domain/hash"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the delete adapter
type Adapter interface {
	ToBytes(ins Delete) ([]byte, error)
	ToInstance(bytes []byte) (Delete, error)
}

// Builder represents a delete builder
type Builder interface {
	Create() Builder
	WithList(list string) Builder
	WithIndex(index string) Builder
	Now() (Delete, error)
}

// Delete represents a delete
type Delete interface {
	Hash() hash.Hash
	List() string
	Idx() string
}
