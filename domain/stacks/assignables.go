package stacks

type assignables struct {
	list []Assignable
}

func createAssignables(
	list []Assignable,
) Assignables {
	out := assignables{
		list: list,
	}

	return &out
}

// List returns the assignables
func (obj *assignables) List() []Assignable {
	return obj.list
}
