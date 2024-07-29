package states

import "github.com/steve-care-software/webx/engine/databases/bytes/domain/headers/states/containers"

// NewStatesForTests creates a new states for tests
func NewStatesForTests(list []State) States {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewStateForTests creates a new state for tests
func NewStateForTests(containers containers.Containers, isDeleted bool) State {
	builder := NewStateBuilder().Create().WithContainers(containers)
	if isDeleted {
		builder.IsDeleted()
	}

	ins, err := builder.Now()
	if err != nil {
		panic(err)
	}

	return ins
}
