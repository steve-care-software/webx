package transactions

import (
	"math/big"

	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"
)

type transaction struct {
	hash   hash.Hash
	asset  hash.Hash
	proof  big.Int
	mine   hash.Hash
	pScore *big.Int
}

func createTransaction(
	hash hash.Hash,
	asset hash.Hash,
	proof big.Int,
	mine hash.Hash,
	pScore *big.Int,
) Transaction {
	out := transaction{
		hash:   hash,
		asset:  asset,
		proof:  proof,
		mine:   mine,
		pScore: pScore,
	}

	return &out
}

// Hash returns the hash
func (obj *transaction) Hash() hash.Hash {
	return obj.hash
}

// Asset returns the asset
func (obj *transaction) Asset() hash.Hash {
	return obj.asset
}

// Proof returns the proof
func (obj *transaction) Proof() big.Int {
	return obj.proof
}

// Mine returns the mine
func (obj *transaction) Mine() hash.Hash {
	return obj.mine
}

// Score returns the score
func (obj *transaction) Score() *big.Int {
	return obj.pScore
}
