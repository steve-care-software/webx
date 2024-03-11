package deletes

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	context     string
	path        string
	identifier  string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		context:     "",
		path:        "",
		identifier:  "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithContext adds a context to the builder
func (app *builder) WithContext(context string) Builder {
	app.context = context
	return app
}

// WithPath adds a path to the builder
func (app *builder) WithPath(path string) Builder {
	app.path = path
	return app
}

// WithIdentifier adds an identifier to the builder
func (app *builder) WithIdentifier(identifier string) Builder {
	app.identifier = identifier
	return app
}

// Now builds a new Delete instance
func (app *builder) Now() (Delete, error) {
	if app.context == "" {
		return nil, errors.New("the context is mandatory in order to build a Delete instance")
	}

	if app.path == "" {
		return nil, errors.New("the path is mandatory in order to build a Delete instance")
	}

	if app.identifier == "" {
		return nil, errors.New("the identifier is mandatory in order to build a Delete instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{})
	if err != nil {
		return nil, err
	}

	return createDelete(*pHash, app.context, app.path, app.identifier), nil
}
