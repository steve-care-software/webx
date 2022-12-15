package everythings

import "github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"

const minEverythingSize = hash.Size * 2

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	hashAdapter := hash.NewAdapter()
	builder := NewBuilder()
	return createAdapter(hashAdapter, builder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents an everything adapter
type Adapter interface {
	ToContent(ins Everything) ([]byte, error)
	ToEverything(content []byte) (Everything, error)
}

// Builder represents an everything builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithException(exception hash.Hash) Builder
	WithEscape(escape hash.Hash) Builder
	Now() (Everything, error)
}

// Everything represents an everything
type Everything interface {
	Hash() hash.Hash
	Exception() hash.Hash
	HasEscape() bool
	Escape() *hash.Hash
}
