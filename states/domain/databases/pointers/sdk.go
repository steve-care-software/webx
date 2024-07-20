package pointers

import (
	"github.com/steve-care-software/datastencil/states/domain/databases/metadatas"
	"github.com/steve-care-software/datastencil/states/domain/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents a pointer adapter
type Adapter interface {
	ToBytes(ins Pointer) ([]byte, error)
	ToInstance(bytes []byte) (Pointer, error)
}

// Builder represents a pointer builder
type Builder interface {
	Create() Builder
	WithHead(head hash.Hash) Builder
	WithMetaData(metaData metadatas.MetaData) Builder
	Now() (Pointer, error)
}

// Pointer represents the database pointer
type Pointer interface {
	Hash() hash.Hash
	Head() hash.Hash
	MetaData() metadatas.MetaData
}
