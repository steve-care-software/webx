package actions

import "github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/values"

// NewActionsForTests creates a new actions for tests
func NewActionsForTests(list []Action) Actions {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewActionWithDeleteForTests creates a new action with delete for tests
func NewActionWithDeleteForTests(path []string) Action {
	ins, err := NewActionBuilder().Create().IsDelete().WithPath(path).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewActionWithInsertForTests creates a new action with insert for tests
func NewActionWithInsertForTests(path []string, insert values.Value) Action {
	ins, err := NewActionBuilder().Create().WithInsert(insert).WithPath(path).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
