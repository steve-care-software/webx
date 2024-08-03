package states

import (
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states/pointers"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states/pointers/delimiters"
)

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

// NewStateWithRootForTests creates a new state with root for tests
func NewStateWithRootForTests(root delimiters.Delimiter, isDeleted bool) State {
	builder := NewStateBuilder().Create().WithRoot(root)
	if isDeleted {
		builder.IsDeleted()
	}

	ins, err := builder.Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewStateWithPointersForTests creates a new state with pointers for tests
func NewStateWithPointersForTests(pointers pointers.Pointers, isDeleted bool) State {
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

// NewStateWithRootAndPointersForTests creates a new state with root and pointers for tests
func NewStateWithRootAndPointersForTests(root delimiters.Delimiter, pointers pointers.Pointers, isDeleted bool) State {
	builder := NewStateBuilder().Create().WithRoot(root).WithPointers(pointers)
	if isDeleted {
		builder.IsDeleted()
	}

	ins, err := builder.Now()
	if err != nil {
		panic(err)
	}

	return ins
}
