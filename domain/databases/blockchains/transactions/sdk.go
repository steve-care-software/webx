package transactions

import (
	"math/big"

	"github.com/steve-care-software/webx/domain/cryptography/hash"
	"github.com/steve-care-software/webx/domain/databases/entities"
)

// Builder represents a transaction builder
type Builder interface {
	Create() Builder
	WithEntity(entity entities.Entity) Builder
	WithOrigin(origin hash.Hash) Builder
	WithAsset(asset hash.Hash) Builder
	WithProof(proof big.Int) Builder
	Now() (Transaction, error)
}

// Transaction represents a transaction
type Transaction interface {
	Entity() entities.Entity
	Origin() hash.Hash
	Asset() hash.Hash
	Proof() big.Int
}
