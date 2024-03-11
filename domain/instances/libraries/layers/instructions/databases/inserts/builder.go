package inserts

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	context     string
	instance    string
	path        string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		context:     "",
		instance:    "",
		path:        "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithContext adds a context to the builder
func (app *builder) WithContext(context string) Builder {
	app.context = context
	return app
}

// WithInstance adds an instance to the builder
func (app *builder) WithInstance(instance string) Builder {
	app.instance = instance
	return app
}

// WithPath adds a path to the builder
func (app *builder) WithPath(path string) Builder {
	app.path = path
	return app
}

// Now builds a new Insert instance
func (app *builder) Now() (Insert, error) {
	if app.context == "" {
		return nil, errors.New("the context is mandatory in order to build an Insert instance")
	}

	if app.instance == "" {
		return nil, errors.New("the instance is mandatory in order to build an Insert instance")
	}

	if app.path == "" {
		return nil, errors.New("the path is mandatory in order to build an Insert instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.context),
		[]byte(app.instance),
		[]byte(app.path),
	})

	if err != nil {
		return nil, err
	}

	return createInsert(*pHash, app.context, app.instance, app.path), nil
}
