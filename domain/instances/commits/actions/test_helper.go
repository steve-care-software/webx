package actions

import (
	"github.com/steve-care-software/datastencil/domain/instances/commits/actions/pointers"
	"github.com/steve-care-software/datastencil/domain/instances/commits/actions/resources"
)

// NewActionsForTests creates a new actions for tests
func NewActionsForTests(list []Action) Actions {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewActionWithInsertForTests creates an action with insert for tests
func NewActionWithInsertForTests(insert resources.Resource) Action {
	ins, err := NewActionBuilder().Create().WithInsert(insert).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewActionWithDeleteForTests creates an action with delete for tests
func NewActionWithDeleteForTests(del pointers.Pointer) Action {
	ins, err := NewActionBuilder().Create().WithDelete(del).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewActionWithInsertAndDeleteForTests creates an action with insert delete for tests
func NewActionWithInsertAndDeleteForTests(insert resources.Resource, del pointers.Pointer) Action {
	ins, err := NewActionBuilder().Create().WithInsert(insert).WithDelete(del).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
