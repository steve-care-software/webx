package deletes

import "github.com/steve-care-software/datastencil/domain/hash"

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
	WithContext(context string) Builder
	WithPath(path string) Builder
	WithIdentifier(identifier string) Builder
	Now() (Delete, error)
}

// Delete represents a delete
type Delete interface {
	Hash() hash.Hash
	Context() string
	Path() string
	Identifier() string
}
