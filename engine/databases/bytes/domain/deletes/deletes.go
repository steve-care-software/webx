package deletes

type deletes struct {
	list []Delete
}

func createDeletes(
	list []Delete,
) Deletes {
	out := deletes{
		list: list,
	}

	return &out
}

// List returns the list of deletes
func (obj *deletes) List() []Delete {
	return obj.list
}
