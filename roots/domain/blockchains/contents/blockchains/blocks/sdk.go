package blocks

import (
	"math/big"

	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"
	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hashtrees"
)

const minBlockSize = hash.Size + 8*2
const errorStr = "the content was expected to contain at least %d bytes in order to convert data to a Transaction instance, %d provided"

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	hashAdapter := hash.NewAdapter()
	hashTreeAdapter := hashtrees.NewAdapter()
	builder := NewBuilder()
	return createAdapter(hashAdapter, hashTreeAdapter, builder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents the block adapter
type Adapter interface {
	ToContent(ins Block) ([]byte, error)
	ToBlock(content []byte) (Block, error)
}

// Builder represents  block builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
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
	NextScore() *big.Int
	PendingScore() *big.Int
	Transactions() hashtrees.HashTree
	HasPrevious() bool
	Previous() *hash.Hash
}
