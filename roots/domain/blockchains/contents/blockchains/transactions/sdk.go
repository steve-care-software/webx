package transactions

import (
	"math/big"

	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"
)

const minTransactionSize = hash.Size*2 + 1

// NewAdapter creates a new adapter
func NewAdapter() Adapter {
	hashAdapter := hash.NewAdapter()
	builder := NewBuilder()
	return createAdapter(hashAdapter, builder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents the transaction adapter
type Adapter interface {
	ToContent(ins Transaction) ([]byte, error)
	ToTransaction(content []byte) (Transaction, error)
}

// Builder represents a transaction builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithAsset(asset hash.Hash) Builder
	WithProof(proof big.Int) Builder
	Now() (Transaction, error)
}

// Transaction represents a transaction
type Transaction interface {
	Hash() hash.Hash
	Asset() hash.Hash
	Proof() *big.Int
}
