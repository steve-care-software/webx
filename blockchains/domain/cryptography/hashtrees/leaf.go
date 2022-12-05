package hashtrees

import (
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
)

type leaf struct {
	head   hash.Hash
	parent ParentLeaf
}

func createLeaf(head hash.Hash) Leaf {
	out := leaf{
		head:   head,
		parent: nil,
	}

	return &out
}

func createLeafWithParent(head hash.Hash, parent ParentLeaf) Leaf {
	out := leaf{
		head:   head,
		parent: parent,
	}

	return &out
}

// Head returns the head hash
func (obj *leaf) Head() hash.Hash {
	return obj.head
}

// HasParent returns true if there is a parent, false otherwise
func (obj *leaf) HasParent() bool {
	return obj.parent != nil
}

// Parent returns the parent, if any
func (obj *leaf) Parent() ParentLeaf {
	return obj.parent
}

// Height returns the leaf height
func (obj *leaf) Height() uint {
	cpt := uint(0)
	var oneLeaf Leaf
	for {

		if oneLeaf == nil {
			oneLeaf = obj
		}

		if !oneLeaf.HasParent() {
			return cpt
		}

		cpt++
		oneLeaf = oneLeaf.Parent().Left()
	}
}
