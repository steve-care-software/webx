package singles

import (
	"errors"

	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/storages"
)

type builder struct {
	storage storages.Storage
	bytes   []byte
}

func createBuilder() Builder {
	out := builder{
		storage: nil,
		bytes:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithStorage adds storage to the builder
func (app *builder) WithStorage(storage storages.Storage) Builder {
	app.storage = storage
	return app
}

// WithBytes adds bytes to the builder
func (app *builder) WithBytes(bytes []byte) Builder {
	app.bytes = bytes
	return app
}

// Now builds a new Single instance
func (app *builder) Now() (Single, error) {
	if app.storage == nil {
		return nil, errors.New("the storage is mandatory in order to build a Single instance")
	}

	if app.bytes != nil && len(app.bytes) <= 0 {
		app.bytes = nil
	}

	if app.bytes == nil {
		return nil, errors.New("the bytes is mandatory in order to build a Single instance")
	}

	return createSingle(
		app.storage,
		app.bytes,
	), nil
}
