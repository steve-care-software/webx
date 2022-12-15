package hashtrees

type leaves struct {
	list []Leaf
}

func createLeaves(list []Leaf) Leaves {
	out := leaves{
		list: list,
	}

	return &out
}

// Leaves returns the leaves
func (obj *leaves) Leaves() []Leaf {
	out := []Leaf{}
	for _, oneLeaf := range obj.list {
		out = append(out, oneLeaf)
	}

	return out
}

// Merge merge Leaves instances
func (obj *leaves) Merge(lves Leaves) Leaves {
	for _, oneLeaf := range lves.Leaves() {
		obj.list = append(obj.list, oneLeaf.(*leaf))
	}

	return obj
}
