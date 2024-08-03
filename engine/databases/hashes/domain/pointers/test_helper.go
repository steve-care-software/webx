package pointers

import (
	bytes_pointers "github.com/steve-care-software/webx/engine/databases/bytes/domain/states/pointers"
	"github.com/steve-care-software/webx/engine/databases/entities/domain/hash"
)

// NewPointersForTests creates a new pointers for tests
func NewPointersForTests(list []Pointer) Pointers {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewPointerForTests creates a new pointer for tests
func NewPointerForTests(hash hash.Hash, pointer bytes_pointers.Pointer) Pointer {
	ins, err := NewPointerBuilder().Create().WithHash(hash).WithPointer(pointer).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
