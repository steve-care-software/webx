package assignments

import (
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
)

type assignment struct {
	index uint
	value hash.Hash
}

func createAssignment(
	index uint,
	value hash.Hash,
) Assignment {
	out := assignment{
		index: index,
		value: value,
	}

	return &out
}

// Index returns the index
func (obj *assignment) Index() uint {
	return obj.index
}

// Value returns the value
func (obj *assignment) Value() hash.Hash {
	return obj.value
}
