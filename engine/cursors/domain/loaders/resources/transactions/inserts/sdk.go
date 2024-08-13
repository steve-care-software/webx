package inserts

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents an insert builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithBytes(bytes []byte) Builder
	WithWhitelist(whitelist []hash.Hash) Builder
	Now() (Insert, error)
}

// Insert represents an insert
type Insert interface {
	Hash() hash.Hash
	Name() string
	Bytes() []byte
	Whitelist() []hash.Hash
}
