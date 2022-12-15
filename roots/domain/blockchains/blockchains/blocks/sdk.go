package blocks

import (
	"math/big"

	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"
	"github.com/steve-care-software/webx/roots/domain/blockchains/blockchains/transactions"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Builder represents a block builder
type Builder interface {
	Create() Builder
	WithHeight(height uint) Builder
	WithPrevious(previous hash.Hash) Builder
	WithTransactions(trx transactions.Transactions) Builder
	WithNextScore(nextScore big.Int) Builder
	WithPendingScore(pendingScore big.Int) Builder
	Now() (Block, error)
}

// Block represents a block
type Block interface {
	Hash() hash.Hash
	Height() uint
	Score() Score
	IsStuck() bool
	Transactions() transactions.Transactions
	HasPrevious() bool
	Previous() *hash.Hash
}

// Score represents a score
type Score interface {
	Next() *big.Int
	Pending() *big.Int
	Total() *big.Int
}
