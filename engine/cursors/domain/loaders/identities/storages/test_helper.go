package storages

import "github.com/steve-care-software/webx/engine/cursors/domain/loaders/namespaces/singles/versions/singles/workspaces/singles/branches/singles/states/singles/pointers/storages"

// NewStoragesForTests creates a new storages for tests
func NewStoragesForTests(list []Storage) Storages {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewStorageForTests creates a new storage for tests
func NewStorageForTests(name string, pointer storages.Storage) Storage {
	ins, err := NewStorageBuilder().Create().WithName(name).WithPointer(pointer).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
