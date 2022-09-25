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

// List returns the trees
func (obj *trees) List() []Tree {
	return obj.list
}
