package transactions

import (
	"math/big"

	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
)

type transaction struct {
	hash  hash.Hash
	asset hash.Hash
	proof *big.Int
}

func createTransaction(
	hash hash.Hash,
	asset hash.Hash,
	proof *big.Int,
) Transaction {
	out := transaction{
		hash:  hash,
		asset: asset,
		proof: proof,
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
func (obj *transaction) Proof() *big.Int {
	return obj.proof
}
