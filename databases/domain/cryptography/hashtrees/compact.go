package hashtrees

import (
	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

type compact struct {
	head   hash.Hash
	leaves Leaves
}

func createCompact(head hash.Hash, leaves Leaves) Compact {
	out := compact{
		head:   head,
		leaves: leaves,
	}

	return &out
}

// Head returns the head hash
func (obj *compact) Head() hash.Hash {
	return obj.head
}

// Leaves returns the leaves
func (obj *compact) Leaves() Leaves {
	return obj.leaves
}

// Length returns the length of the compact hashtree
func (obj *compact) Length() uint {
	return uint(len(obj.leaves.Leaves()))
}
