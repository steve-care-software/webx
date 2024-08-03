package pointers

import (
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states/pointers/delimiters"
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
func NewPointerForTests(hash hash.Hash, delimiter delimiters.Delimiter) Pointer {
	ins, err := NewPointerBuilder().Create().WithHash(hash).WithDelimiter(delimiter).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
