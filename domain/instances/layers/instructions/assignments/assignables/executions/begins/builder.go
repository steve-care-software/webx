package begins

import (
	"errors"

	"github.com/steve-care-software/historydb/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	path        string
	context     string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		path:        "",
		context:     "",
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

// WithContext adds a context to the builder
func (app *builder) WithContext(context string) Builder {
	app.context = context
	return app
}

// Now builds a new Begin instance
func (app *builder) Now() (Begin, error) {
	if app.path == "" {
		return nil, errors.New("the path is mandatory in order to build a Begin instance")
	}

	if app.context == "" {
		return nil, errors.New("the context is mandatory in order to build a Begin instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.path),
		[]byte(app.context),
	})

	if err != nil {
		return nil, err
	}

	return createBegin(*pHash, app.path, app.context), nil
}
