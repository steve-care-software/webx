package inserts

import "github.com/steve-care-software/webx/engine/cursors/domain/hash"

// NewInsertForTests creates a new insert for tests
func NewInsertForTests(name string, bytes []byte, whitelist []hash.Hash) Insert {
	ins, err := NewBuilder().Create().WithName(name).WithBytes(bytes).WithWhitelist(whitelist).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
