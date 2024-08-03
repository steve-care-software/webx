package instructions

import (
	"github.com/steve-care-software/webx/engine/hashes/domain/hash"
)

type instructions struct {
	hash hash.Hash
	list []Instruction
}

func createInstructions(
	hash hash.Hash,
	list []Instruction,
) Instructions {
	out := instructions{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *instructions) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *instructions) List() []Instruction {
	return obj.list
}
