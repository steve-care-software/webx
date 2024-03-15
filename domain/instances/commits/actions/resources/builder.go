package resources

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances"
)

type builder struct {
	hashAdapter hash.Adapter
	path        []string
	instance    instances.Instance
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		path:        nil,
		instance:    nil,
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
func (app *builder) WithPath(path []string) Builder {
	app.path = path
	return app
}

// WithInstance adds an instance to the builder
func (app *builder) WithInstance(instance instances.Instance) Builder {
	app.instance = instance
	return app
}

// Now builds a new Resource instance
func (app *builder) Now() (Resource, error) {
	if app.path != nil && len(app.path) <= 0 {
		app.path = nil
	}

	if app.path == nil {
		return nil, errors.New("the path is mandatory in order to build a Resource instance")
	}

	if app.instance == nil {
		return nil, errors.New("the instance is mandatory in order to build a Resource instance")
	}

	data := [][]byte{
		app.instance.Hash().Bytes(),
	}

	for _, onePath := range app.path {
		data = append(data, []byte(onePath))
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createResource(*pHash, app.path, app.instance), nil
}
