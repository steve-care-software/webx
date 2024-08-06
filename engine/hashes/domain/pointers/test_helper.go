package pointers

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/delimiters"
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
func NewPointerForTests(hash hash.Hash, delimiter delimiters.Delimiter) Pointer {
	ins, err := NewPointerBuilder().Create().WithHash(hash).WithDelimiter(delimiter).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
