package pointers

import "github.com/steve-care-software/webx/engine/cursors/domain/loaders/namespaces/singles/versions/singles/workspaces/singles/branches/singles/states/singles/pointers/storages"

// NewPointersForTests creates a new pointers for tests
func NewPointersForTests(list []Pointer) Pointers {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewPointerForTests creates a new pointer for tests
func NewPointerForTests(storage storages.Storage, bytes []byte) Pointer {
	ins, err := NewPointerBuilder().Create().WithStorage(storage).WithBytes(bytes).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
