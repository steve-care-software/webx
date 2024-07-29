package pointers

import "github.com/steve-care-software/webx/engine/databases/bytes/domain/retrievals"

// NewPointersForTests creates a new pointers for tests
func NewPointersForTests(list []Pointer) Pointers {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewPointerForTests creates a new pointer for tests
func NewPointerForTests(retireval retrievals.Retrieval, isDeleted bool) Pointer {
	builder := NewPointerBuilder().Create().WithRetrieval(retireval)
	if isDeleted {
		builder.IsDeleted()
	}

	ins, err := builder.Now()
	if err != nil {
		panic(err)
	}

	return ins
}
