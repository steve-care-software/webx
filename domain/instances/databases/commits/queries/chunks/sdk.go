package chunks

import (
	"github.com/steve-care-software/datastencil/domain/hash"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents a chunk builder
type Builder interface {
	Create() Builder
	WithPath(path []string) Builder
	WithFingerprint(fingerprint hash.Hash) Builder
	Now() (Chunk, error)
}

// Chunk represents a chunk
type Chunk interface {
	Hash() hash.Hash
	Path() []string
	Fingerprint() hash.Hash
}

// Repository represents a chunk repository
type Repository interface {
	Retrieve(hash hash.Hash) (Chunk, error)
}

// Service represents a chunk service
type Service interface {
	Save(ins Chunk) error
	Delete(hash hash.Hash) error
}
