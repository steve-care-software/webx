package storages

import (
	"errors"

	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/namespaces/singles/versions/singles/workspaces/singles/branches/singles/states/singles/pointers/storages"
)

type storageBuilder struct {
	name    string
	pointer storages.Storage
}

func createStorageBuilder() StorageBuilder {
	out := storageBuilder{
		name:    "",
		pointer: nil,
	}

	return &out
}

// Create initializes the builder
func (app *storageBuilder) Create() StorageBuilder {
	return createStorageBuilder()
}

// WithName adds a name to the builder
func (app *storageBuilder) WithName(name string) StorageBuilder {
	app.name = name
	return app
}

// WithPointer adds a pointer to the builder
func (app *storageBuilder) WithPointer(pointer storages.Storage) StorageBuilder {
	app.pointer = pointer
	return app
}

// Now builds a new Storage instance
func (app *storageBuilder) Now() (Storage, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Storage instance")
	}

	if app.pointer == nil {
		return nil, errors.New("the pointer is mandatory in order to build a Storage instance")
	}

	return createStorage(app.name, app.pointer), nil
}
