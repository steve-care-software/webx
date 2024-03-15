package pointers

import "github.com/steve-care-software/datastencil/domain/hash"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Builder represents a pointer builder
type Builder interface {
	Create() Builder
	WithPath(path []string) Builder
	WithIdentifier(identifier hash.Hash) Builder
	Now() (Pointer, error)
}

// Pointer represents a pointer
type Pointer interface {
	Hash() hash.Hash
	Path() []string
	Identifier() hash.Hash
}
