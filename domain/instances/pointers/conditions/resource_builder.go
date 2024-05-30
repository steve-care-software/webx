package conditions

import (
	"errors"
	"path/filepath"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type resourceBuilder struct {
	hashAdapter  hash.Adapter
	path         []string
	mustBeLoaded bool
}

func createResourceBuilder(
	hashAdapter hash.Adapter,
) ResourceBuilder {
	out := resourceBuilder{
		hashAdapter:  hashAdapter,
		path:         nil,
		mustBeLoaded: false,
	}

	return &out
}

// Create initializes the builder
func (app *resourceBuilder) Create() ResourceBuilder {
	return createResourceBuilder(
		app.hashAdapter,
	)
}

// WithPath adds a path to the builder
func (app *resourceBuilder) WithPath(path []string) ResourceBuilder {
	app.path = path
	return app
}

// MustBeLoaded flags the builder as must be loaded
func (app *resourceBuilder) MustBeLoaded() ResourceBuilder {
	app.mustBeLoaded = true
	return app
}

// Now builds a new Resource instance
func (app *resourceBuilder) Now() (Resource, error) {
	if app.path != nil && len(app.path) <= 0 {
		app.path = nil
	}

	if app.path == nil {
		return nil, errors.New("the path is mandatory in order to build a Resource instance")
	}

	mustBeLoaded := "false"
	if app.mustBeLoaded {
		mustBeLoaded = "true"
	}

	path := filepath.Join(app.path...)
	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(mustBeLoaded),
		[]byte(path),
	})

	if err != nil {
		return nil, err
	}

	return createResource(*pHash, app.path, app.mustBeLoaded), nil
}
