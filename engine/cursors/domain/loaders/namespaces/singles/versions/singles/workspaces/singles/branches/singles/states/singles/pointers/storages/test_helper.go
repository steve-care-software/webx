package storages

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/delimiters"
)

// NewStoragesForTests creates a new storages for tests
func NewStoragesForTests(list []Storage) Storages {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewStorageForTests creates a new storage for tests
func NewStorageForTests(delimiter delimiters.Delimiter, isDeleted bool) Storage {
	builder := NewStorageBuilder().Create().WithDelimiter(delimiter)
	if isDeleted {
		builder.IsDeleted()
	}

	ins, err := builder.Now()
	if err != nil {
		panic(err)
	}

	return ins
}
