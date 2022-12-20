package commits

import (
	"math/big"

	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

type mine struct {
	result     hash.Hash
	pProof     *big.Int
	difficulty uint
}

func createMine(
	result hash.Hash,
	pProof *big.Int,
	difficulty uint,
) Mine {
	out := mine{
		result:     result,
		pProof:     pProof,
		difficulty: difficulty,
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

// Difficulty returns the difficulty
func (obj *mine) Difficulty() uint {
	return obj.difficulty
}
