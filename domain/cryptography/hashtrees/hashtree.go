package hashtrees

import (
	"github.com/steve-care-software/webx/domain/cryptography/hash"
)

// HTree represents a concrete HashTree implementation
type hashtree struct {
	head   hash.Hash
	parent ParentLeaf
}

func createHashTree(head hash.Hash, parent ParentLeaf) HashTree {
	out := hashtree{
		head:   head,
		parent: parent,
	}

	return &out
}

// Height returns the hashtree height
func (obj *hashtree) Height() uint {
	left := obj.parent.Left()
	return left.Height() + 2
}

// Head returns the head hash
func (obj *hashtree) Head() hash.Hash {
	return obj.head
}

// Parent returns the parent leaf
func (obj *hashtree) Parent() ParentLeaf {
	return obj.parent
}
