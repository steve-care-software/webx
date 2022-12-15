package programs

import "github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"

type program struct {
	hash         hash.Hash
	instructions []hash.Hash
	outputs      []uint
}

func createProgram(
	hash hash.Hash,
	instructions []hash.Hash,
) Program {
	return createProgramInternally(hash, instructions, nil)
}

func createProgramWithOutputs(
	hash hash.Hash,
	instructions []hash.Hash,
	outputs []uint,
) Program {
	return createProgramInternally(hash, instructions, outputs)
}

func createProgramInternally(
	hash hash.Hash,
	instructions []hash.Hash,
	outputs []uint,
) Program {
	out := program{
		hash:         hash,
		instructions: instructions,
		outputs:      outputs,
	}

	return &out
}

// Hash returns the hash
func (obj *program) Hash() hash.Hash {
	return obj.hash
}

// Instructions returns the instructions
func (obj *program) Instructions() []hash.Hash {
	return obj.instructions
}

// HasOutputs returns true if there is outputs, false otherwise
func (obj *program) HasOutputs() bool {
	return obj.outputs != nil
}

// Outputs returns the outputs, if any
func (obj *program) Outputs() []uint {
	return obj.outputs
}
