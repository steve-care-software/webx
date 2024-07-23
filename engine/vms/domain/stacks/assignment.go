package stacks

type assignment struct {
	name       string
	assignable Assignable
}

func createAssignment(
	name string,
	assignable Assignable,
) Assignment {
	out := assignment{
		name:       name,
		assignable: assignable,
	}

	return &out
}

// Name returns the name
func (obj *assignment) Name() string {
	return obj.name
}

// Assignable returns the assignable
func (obj *assignment) Assignable() Assignable {
	return obj.assignable
}
