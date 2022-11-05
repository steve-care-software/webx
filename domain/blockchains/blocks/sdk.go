package blocks

import (
	"math/big"

	"github.com/steve-care-software/webx/domain/cryptography/hash"
	"github.com/steve-care-software/webx/domain/blockchains/transactions"
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

// Repository represents a block repository
type Repository interface {
	RetrieveByHeight(reference hash.Hash, height uint) (Block, error)
	RetrieveByHash(reference hash.Hash, hash hash.Hash) (Block, error)
	RetrieveByPreviousHash(reference hash.Hash, prev hash.Hash) (Block, error)
}

// Service represents a block service
type Service interface {
	Save(reference hash.Hash, block Block) error
	Delete(reference hash.Hash, block Block) error
}
