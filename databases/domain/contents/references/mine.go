package references

import (
	"math/big"

	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

type mine struct {
	result hash.Hash
	pProof *big.Int
	score  uint
}

func createMine(
	result hash.Hash,
	pProof *big.Int,
	score uint,
) Mine {
	out := mine{
		result: result,
		pProof: pProof,
		score:  score,
	}

	return &out
}

// Result returns the result
func (obj *mine) Result() hash.Hash {
	return obj.result
}

// Proof returns the proof
func (obj *mine) Proof() *big.Int {
	return obj.pProof
}

// Score returns the score
func (obj *mine) Score() uint {
	return obj.score
}
