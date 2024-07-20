package retrieves

import "github.com/steve-care-software/datastencil/states/domain/hash"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the retrieve adapter
type Adapter interface {
	ToBytes(ins Retrieve) ([]byte, error)
	ToInstance(bytes []byte) (Retrieve, error)
}

// Builder represents a retrieve builder
type Builder interface {
	Create() Builder
	WithContext(context string) Builder
	WithIndex(index string) Builder
	WithLength(length string) Builder
	Now() (Retrieve, error)
}

// Retrieve represents a retrieve instruction
type Retrieve interface {
	Hash() hash.Hash
	Context() string
	Index() string
	HasLength() bool
	Length() string
}
