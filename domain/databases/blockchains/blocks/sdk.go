package blocks

import (
	"math/big"

	"github.com/steve-care-software/webx/domain/cryptography/hash"
	"github.com/steve-care-software/webx/domain/databases/entities"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents  block builder
type Builder interface {
	Create() Builder
	WithEntity(entity entities.Entity) Builder
	WithHeight(height uint) Builder
	WithNextScore(nextScore big.Int) Builder
	WithPendingScope(pendingScore big.Int) Builder
	WithTransactions(transactions entities.Identifiers) Builder
	WithPrevious(previous hash.Hash) Builder
	Now() (Block, error)
}

// Block represents a block
type Block interface {
	Entity() entities.Entity
	Height() uint
	NextScore() big.Int
	PendingScore() big.Int
	Transactions() entities.Identifiers
	HasPrevious() bool
	Previous() *hash.Hash
}
