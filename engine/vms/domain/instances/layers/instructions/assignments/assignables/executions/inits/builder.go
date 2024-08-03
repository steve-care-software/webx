package inits

import (
	"errors"

	"github.com/steve-care-software/webx/engine/hashes/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	path        string
	name        string
	description string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		path:        "",
		name:        "",
		description: "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithPath adds a path to the builder
func (app *builder) WithPath(path string) Builder {
	app.path = path
	return app
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithDescription adds a description to the builder
func (app *builder) WithDescription(description string) Builder {
	app.description = description
	return app
}

// Now builds a new Init instance
func (app *builder) Now() (Init, error) {
	if app.path == "" {
		return nil, errors.New("the path is mandatory in order to build an Init instance")
	}

	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build an Init instance")
	}

	if app.description == "" {
		return nil, errors.New("the description is mandatory in order to build an Init instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.path),
		[]byte(app.name),
		[]byte(app.description),
	})

	if err != nil {
		return nil, err
	}

	return createInit(*pHash, app.path, app.name, app.description), nil
}
