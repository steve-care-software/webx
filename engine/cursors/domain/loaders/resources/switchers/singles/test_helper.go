package singles

import "github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/storages"

// NewSingleForTests creates a new single for tests
func NewSingleForTests(storage storages.Storage, bytes []byte) Single {
	ins, err := NewBuilder().Create().WithStorage(storage).WithBytes(bytes).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
