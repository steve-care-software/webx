package pointers

import (
	"errors"

	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/namespaces/singles/versions/singles/workspaces/singles/branches/singles/states/singles/pointers/storages"
)

type pointerBuilder struct {
	storage storages.Storage
	bytes   []byte
}

func createPointerBuilder() PointerBuilder {
	out := pointerBuilder{
		storage: nil,
		bytes:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *pointerBuilder) Create() PointerBuilder {
	return createPointerBuilder()
}

// WithStorage adds a storage to the builder
func (app *pointerBuilder) WithStorage(storage storages.Storage) PointerBuilder {
	app.storage = storage
	return app
}

// WithBytes add bytes to the builder
func (app *pointerBuilder) WithBytes(bytes []byte) PointerBuilder {
	app.bytes = bytes
	return app
}

// Now builds a new Pointer instance
func (app *pointerBuilder) Now() (Pointer, error) {
	if app.storage == nil {
		return nil, errors.New("the storage pointer is mandatory in order to build a Pointer instance")
	}

	if app.bytes != nil && len(app.bytes) <= 0 {
		app.bytes = nil
	}

	if app.bytes == nil {
		return nil, errors.New("the bytes is mandatory in order to build a Pointer instance")
	}

	return createPointer(
		app.storage,
		app.bytes,
	), nil
}
