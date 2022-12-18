package hashtrees

type parentLeaf struct {
	left  Leaf
	right Leaf
}

func createParentLeaf(left Leaf, right Leaf) ParentLeaf {
	out := parentLeaf{
		left:  left,
		right: right,
	}

	return &out
}

// Left returns the left leaf
func (obj *parentLeaf) Left() Leaf {
	return obj.left
}

// Right returns the right leaf
func (obj *parentLeaf) Right() Leaf {
	return obj.right
}
