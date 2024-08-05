package states

import "github.com/steve-care-software/webx/engine/bytes/domain/pointers"

// NewStatesForTests creates a new states for tests
func NewStatesForTests(list []State) States {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewStateForTests creates a new state for tests
func NewStateForTests(isDeleted bool) State {
	builder := NewStateBuilder().Create()
	if isDeleted {
		builder.IsDeleted()
	}

	ins, err := builder.Now()
	if err != nil {
		panic(err)
	}

	return ins
}

func NewStateWithPointersForTests(isDeleted bool, pointers pointers.Pointers) State {
	builder := NewStateBuilder().Create().WithPointers(pointers)
	if isDeleted {
		builder.IsDeleted()
	}

	ins, err := builder.Now()
	if err != nil {
		panic(err)
	}

	return ins
}
