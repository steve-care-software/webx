package blocks

import (
	"math/big"

	"github.com/steve-care-software/webx/domain/cryptography/hash"
	"github.com/steve-care-software/webx/domain/databases/entities"
)

type block struct {
	entity       entities.Entity
	height       uint
	nextScore    big.Int
	pendingScore big.Int
	trx          entities.Identifiers
	previous     *hash.Hash
}

func createBlock(
	entity entities.Entity,
	height uint,
	nextScore big.Int,
	pendingScore big.Int,
	trx entities.Identifiers,
) Block {
	return createBlockInternally(entity, height, nextScore, pendingScore, trx, nil)
}

func createBlockWithPrevious(
	entity entities.Entity,
	height uint,
	nextScore big.Int,
	pendingScore big.Int,
	trx entities.Identifiers,
	previous *hash.Hash,
) Block {
	return createBlockInternally(entity, height, nextScore, pendingScore, trx, previous)
}

func createBlockInternally(
	entity entities.Entity,
	height uint,
	nextScore big.Int,
	pendingScore big.Int,
	trx entities.Identifiers,
	previous *hash.Hash,
) Block {
	out := block{
		entity:       entity,
		height:       height,
		nextScore:    nextScore,
		pendingScore: pendingScore,
		trx:          trx,
		previous:     previous,
	}

	return &out
}

// Entity returns the entity
func (obj *block) Entity() entities.Entity {
	return obj.entity
}

// Height returns the height
func (obj *block) Height() uint {
	return obj.height
}

// NextScore returns the nextScore
func (obj *block) NextScore() big.Int {
	return obj.nextScore
}

// PendingScore returns the pendingScore
func (obj *block) PendingScore() big.Int {
	return obj.pendingScore
}

// Transactions returns the transactions
func (obj *block) Transactions() entities.Identifiers {
	return obj.trx
}

// HasPrevious returns true if there is a previous hash, false otherwise
func (obj *block) HasPrevious() bool {
	return obj.previous != nil
}

// Previous retruns the previous hash, if any
func (obj *block) Previous() *hash.Hash {
	return obj.previous
}
