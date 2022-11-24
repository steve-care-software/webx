package blocks

import (
	"math/big"

	"github.com/steve-care-software/webx/domain/cryptography/hash"
	"github.com/steve-care-software/webx/domain/cryptography/hashtrees"
)

type block struct {
	hash         hash.Hash
	height       uint
	nextScore    big.Int
	pendingScore big.Int
	trx          hashtrees.HashTree
	previous     *hash.Hash
}

func createBlock(
	hash hash.Hash,
	height uint,
	nextScore big.Int,
	pendingScore big.Int,
	trx hashtrees.HashTree,
) Block {
	return createBlockInternally(hash, height, nextScore, pendingScore, trx, nil)
}

func createBlockWithPrevious(
	hash hash.Hash,
	height uint,
	nextScore big.Int,
	pendingScore big.Int,
	trx hashtrees.HashTree,
	previous *hash.Hash,
) Block {
	return createBlockInternally(hash, height, nextScore, pendingScore, trx, previous)
}

func createBlockInternally(
	hash hash.Hash,
	height uint,
	nextScore big.Int,
	pendingScore big.Int,
	trx hashtrees.HashTree,
	previous *hash.Hash,
) Block {
	out := block{
		hash:         hash,
		height:       height,
		nextScore:    nextScore,
		pendingScore: pendingScore,
		trx:          trx,
		previous:     previous,
	}

	return &out
}

// Hash returns the hash
func (obj *block) Hash() hash.Hash {
	return obj.hash
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
func (obj *block) Transactions() hashtrees.HashTree {
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
