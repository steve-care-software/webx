package inputs

import (
	"errors"

	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	value       string
	path        string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		value:       "",
		path:        "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithValue adds a value to the builder
func (app *builder) WithValue(value string) Builder {
	app.value = value
	return app
}

// WithPath adds a path to the builder
func (app *builder) WithPath(path string) Builder {
	app.path = path
	return app
}

// Now builds a new Input instance
func (app *builder) Now() (Input, error) {
	bytes := [][]byte{}
	if app.value != "" {
		bytes = append(bytes, []byte("value"))
		bytes = append(bytes, []byte(app.value))
	}

	if app.path != "" {
		bytes = append(bytes, []byte("path"))
		bytes = append(bytes, []byte(app.path))
	}

	if len(bytes) != 2 {
		return nil, errors.New("the value or path is mandatory in order to build an Input instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(bytes)
	if err != nil {
		return nil, err
	}

	if app.value != "" {
		return createInputWithValue(*pHash, app.value), nil
	}

	return createInputWithPath(*pHash, app.path), nil
}
