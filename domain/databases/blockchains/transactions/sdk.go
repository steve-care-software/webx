package transactions

import (
	"math/big"

	"github.com/steve-care-software/webx/domain/cryptography/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Adapter represents the transaction adapter
type Adapter interface {
	ToContent(ins Transaction) ([]byte, error)
	ToTransaction(content []byte) (Transaction, error)
}

// Builder represents a transaction builder
type Builder interface {
	Create() Builder
	WithAsset(asset hash.Hash) Builder
	WithProof(proof big.Int) Builder
	Now() (Transaction, error)
}

// Transaction represents a transaction
type Transaction interface {
	Hash() hash.Hash
	Asset() hash.Hash
	Proof() big.Int
}
