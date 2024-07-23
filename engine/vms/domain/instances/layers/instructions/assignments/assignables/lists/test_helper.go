package lists

import "github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/lists/fetches"

// NewListWithFetchForTests creates a new list with fetch for tests
func NewListWithFetchForTests(fetch fetches.Fetch) List {
	ins, err := NewBuilder().Create().WithFetch(fetch).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewListWithLengthForTests creates a new list with length for tests
func NewListWithLengthForTests(length string) List {
	ins, err := NewBuilder().Create().WithLength(length).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewListWithCreateForTests creates a new list with create for tests
func NewListWithCreateForTests(create string) List {
	ins, err := NewBuilder().Create().WithCreate(create).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
