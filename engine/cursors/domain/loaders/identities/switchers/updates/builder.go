package updates

import (
	"errors"

	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles"
)

type builder struct {
	single singles.Single
	bytes  []byte
}

func createBuilder() Builder {
	out := builder{
		single: nil,
		bytes:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithSingle adds a single to the builder
func (app *builder) WithSingle(single singles.Single) Builder {
	app.single = single
	return app
}

// WithBytes add bytes to the builder
func (app *builder) WithBytes(bytes []byte) Builder {
	app.bytes = bytes
	return app
}

// Now builds a new Update instance
func (app *builder) Now() (Update, error) {
	if app.single == nil {
		return nil, errors.New("the single identity is mandatory in order to build an Update instance")
	}

	if app.bytes != nil && len(app.bytes) <= 0 {
		app.bytes = nil
	}

	if app.bytes == nil {
		return nil, errors.New("the bytes are mandatory in order to build an Update instance")
	}

	return createUpdate(
		app.single,
		app.bytes,
	), nil
}
