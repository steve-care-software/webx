package storages

import (
	"errors"

	"github.com/steve-care-software/webx/engine/cursors/domain/storages/delimiters"
)

type storageBuilder struct {
	delimiter delimiters.Delimiter
	isDeleted bool
}

func createStorageBuilder() StorageBuilder {
	out := storageBuilder{
		delimiter: nil,
		isDeleted: false,
	}

	return &out
}

// Create initializes the builder
func (app *storageBuilder) Create() StorageBuilder {
	return createStorageBuilder()
}

// WithDelimiter adds a delimiter to the builder
func (app *storageBuilder) WithDelimiter(delimiter delimiters.Delimiter) StorageBuilder {
	app.delimiter = delimiter
	return app
}

// IsDeleted flags the builder as deleted
func (app *storageBuilder) IsDeleted() StorageBuilder {
	app.isDeleted = true
	return app
}

// Now builds a new Storage instance
func (app *storageBuilder) Now() (Storage, error) {
	if app.delimiter == nil {
		return nil, errors.New("the delimiter is mandatory in order to build a Storage instance")
	}

	if app.isDeleted {
		return createStorageWithDeleted(app.delimiter), nil
	}

	return createStorage(app.delimiter), nil
}
