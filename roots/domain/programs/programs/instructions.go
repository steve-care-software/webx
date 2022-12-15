package programs

import "github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"

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

// List returns the instructions
func (obj *instructions) List() []Instruction {
	return obj.list
}
