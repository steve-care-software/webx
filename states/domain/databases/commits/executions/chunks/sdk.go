package chunks

import "github.com/steve-care-software/datastencil/states/domain/hash"

// NewBuilder creates a new chunk builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents the chunk builder
type Builder interface {
	Create() Builder
	WithPath(path []string) Builder
	WithFingerPrint(fingerPrint hash.Hash) Builder
	Now() (Chunk, error)
}

// Chunk represents a chunk
type Chunk interface {
	Hash() hash.Hash
	Path() []string
	FingerPrint() hash.Hash
}
