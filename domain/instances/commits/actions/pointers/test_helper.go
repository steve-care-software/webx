package pointers

import "github.com/steve-care-software/datastencil/domain/hash"

// NewPointerForTests creates a new pointer for tests
func NewPointerForTests(path []string, identifier hash.Hash) Pointer {
	ins, err := NewBuilder().Create().WithPath(path).WithIdentifier(identifier).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
