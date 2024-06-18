package retrieves

import "github.com/steve-care-software/datastencil/domain/hash"

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
	WithExists(exists string) Builder
	WithRetrieve(retrieve string) Builder
	IsList() Builder
	Now() (Retrieve, error)
}

// Retrieve represents a retrieve
type Retrieve interface {
	Hash() hash.Hash
	IsList() bool
	IsExists() bool
	Exists() string
	IsRetrieve() bool
	Retrieve() string
}
