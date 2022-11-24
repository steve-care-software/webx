package blocks

import (
	"math/big"

	"github.com/steve-care-software/webx/domain/cryptography/hash"
	"github.com/steve-care-software/webx/domain/cryptography/hashtrees"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Adapter represents the block adapter
type Adapter interface {
	ToContent(ins Block) ([]byte, error)
	ToTransaction(content []byte) (Block, error)
}

// Builder represents  block builder
type Builder interface {
	Create() Builder
	WithHeight(height uint) Builder
	WithNextScore(nextScore big.Int) Builder
	WithPendingScope(pendingScore big.Int) Builder
	WithTransactions(transactions hashtrees.HashTree) Builder
	WithPrevious(previous hash.Hash) Builder
	Now() (Block, error)
}

// Block represents a block
type Block interface {
	Hash() hash.Hash
	Height() uint
	NextScore() big.Int
	PendingScore() big.Int
	Transactions() hashtrees.HashTree
	HasPrevious() bool
	Previous() *hash.Hash
}
