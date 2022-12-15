package blocks

import (
	"github.com/steve-care-software/webx/domain/blockchains/transactions"
	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"
)

type block struct {
	hash      hash.Hash
	height    uint
	score     Score
	isStuck   bool
	trx       transactions.Transactions
	pPrevious *hash.Hash
}

func createBlock(
	hash hash.Hash,
	height uint,
	score Score,
	isStuck bool,
	trx transactions.Transactions,
) Block {
	return createBlockInternally(hash, height, score, isStuck, trx, nil)
}

func createBlockWithPrevious(
	hash hash.Hash,
	height uint,
	score Score,
	isStuck bool,
	trx transactions.Transactions,
	pPrevious *hash.Hash,
) Block {
	return createBlockInternally(hash, height, score, isStuck, trx, pPrevious)
}

func createBlockInternally(
	hash hash.Hash,
	height uint,
	score Score,
	isStuck bool,
	trx transactions.Transactions,
	pPrevious *hash.Hash,
) Block {
	out := block{
		hash:      hash,
		height:    height,
		score:     score,
		isStuck:   isStuck,
		trx:       trx,
		pPrevious: pPrevious,
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

// Score returns the score
func (obj *block) Score() Score {
	return obj.score
}

// IsStuck returns true if stuck, false otherwise
func (obj *block) IsStuck() bool {
	return obj.isStuck
}

// Transactions returns the transactions
func (obj *block) Transactions() transactions.Transactions {
	return obj.trx
}

// HasPrevious returns true if there is a previous, false otherwise
func (obj *block) HasPrevious() bool {
	return obj.pPrevious != nil
}

// Previous returns the previous, if any
func (obj *block) Previous() *hash.Hash {
	return obj.pPrevious
}
