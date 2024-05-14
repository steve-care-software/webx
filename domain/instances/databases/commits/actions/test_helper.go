package actions

import "github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications"

// NewActionsForTests creates a new actions for tests
func NewActionsForTests(list []Action) Actions {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewActionWithModificationsForTests creates a new action with insert for tests
func NewActionWithModificationsForTests(path []string, modifications modifications.Modifications) Action {
	ins, err := NewActionBuilder().Create().WithModifications(modifications).WithPath(path).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
