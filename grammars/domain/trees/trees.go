package trees

type trees struct {
	list []Tree
}

func createTrees(
	list []Tree,
) Trees {
	out := trees{
		list: list,
	}

	return &out
}

// Bytes returns the trees' bytes
func (obj *trees) Bytes(includeChannels bool) []byte {
	output := []byte{}
	for _, oneTree := range obj.list {
		output = append(output, oneTree.Bytes(includeChannels)...)
	}

	return output
}

// List returns the trees
func (obj *trees) List() []Tree {
	return obj.list
}
