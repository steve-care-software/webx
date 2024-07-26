package instructions

import "github.com/steve-care-software/webx/engine/databases/entities/domain/hash"

type condition struct {
	hash         hash.Hash
	variable     string
	instructions Instructions
}

func createCondition(
	hash hash.Hash,
	variable string,
	instructions Instructions,
) Condition {
	out := condition{
		hash:         hash,
		variable:     variable,
		instructions: instructions,
	}

	return &out
}

// Hash returns the hash
func (obj *condition) Hash() hash.Hash {
	return obj.hash
}

// Variable returns the variable
func (obj *condition) Variable() string {
	return obj.variable
}

// Instructions returns the instructions
func (obj *condition) Instructions() Instructions {
	return obj.instructions
}
