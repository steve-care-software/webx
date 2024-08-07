package pointers

import "github.com/steve-care-software/webx/engine/cursors/domain/storages/pointers"

// NewPointersForTests creates a new pointers for tests
func NewPointersForTests(list []Pointer) Pointers {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewPointerForTests creates a new pointer for tests
func NewPointerForTests(storage pointers.Pointer, bytes []byte) Pointer {
	ins, err := NewPointerBuilder().Create().WithStorage(storage).WithBytes(bytes).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
