package transactions

import (
	"math/big"

	"github.com/steve-care-software/webx/domain/cryptography/hash"
	"github.com/steve-care-software/webx/domain/databases/entities"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a transaction builder
type Builder interface {
	Create() Builder
	WithEntity(entity entities.Entity) Builder
	WithAsset(asset hash.Hash) Builder
	WithProof(proof big.Int) Builder
	Now() (Transaction, error)
}

// Transaction represents a transaction
type Transaction interface {
	Entity() entities.Entity
	Asset() hash.Hash
	Proof() big.Int
}
