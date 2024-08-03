package reads

import "github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents a read builder
type Builder interface {
	Create() Builder
	WithIdentifier(identifier string) Builder
	WithIndex(index string) Builder
	WithLength(length string) Builder
	Now() (Read, error)
}

// Read represents a read
type Read interface {
	Hash() hash.Hash
	Identifier() string
	HasIndex() bool
	Index() string
	HasLength() bool
	Length() string
}
