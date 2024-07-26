package instructions

import "github.com/steve-care-software/webx/engine/databases/entities/domain/hash"

type loop struct {
	hash         hash.Hash
	amount       string
	instructions Instructions
}

func createLoop(
	hash hash.Hash,
	amount string,
	instructions Instructions,
) Loop {
	out := loop{
		hash:         hash,
		amount:       amount,
		instructions: instructions,
	}

	return &out
}

// Hash returns the hash
func (obj *loop) Hash() hash.Hash {
	return obj.hash
}

// Amount returns the amount
func (obj *loop) Amount() string {
	return obj.amount
}

// Instructions returns the instructions
func (obj *loop) Instructions() Instructions {
	return obj.instructions
}
