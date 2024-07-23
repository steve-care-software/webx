package assignments

import "github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables"

// NewAssignmentForTests creates a new assignment for tests
func NewAssignmentForTests(name string, assignable assignables.Assignable) Assignment {
	ins, err := NewBuilder().Create().WithName(name).WithAssignable(assignable).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
