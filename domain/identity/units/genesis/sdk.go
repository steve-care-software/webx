package genesis

import (
	"github.com/steve-care-software/syntax/domain/identity/cryptography/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Builder represents a genesis builder
type Builder interface {
	Create() Builder
	WithSupply(supply uint64) Builder
	WithOwner(owner []hash.Hash) Builder
	Now() (Genesis, error)
}

// Genesis represents a genesis
type Genesis interface {
	Hash() hash.Hash
	Supply() uint64
	Owner() []hash.Hash
}
