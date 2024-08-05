package pointers

import (
	"github.com/steve-care-software/webx/engine/bytes/domain/delimiters"
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
func NewPointerForTests(delimiter delimiters.Delimiter, isDeleted bool) Pointer {
	builder := NewPointerBuilder().Create().WithDelimiter(delimiter)
	if isDeleted {
		builder.IsDeleted()
	}

	ins, err := builder.Now()
	if err != nil {
		panic(err)
	}

	return ins
}
