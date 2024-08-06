package opens

import (
	"errors"

	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	path        string
	permission  string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		path:        "",
		permission:  "",
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

// WithPermission adds a permission to the builder
func (app *builder) WithPermission(permission string) Builder {
	app.permission = permission
	return app
}

// Now builds a new Open instance
func (app *builder) Now() (Open, error) {
	if app.path == "" {
		return nil, errors.New("the path is mandatory in order to build an Open instance")
	}

	if app.permission == "" {
		return nil, errors.New("the permission is mandatory in order to build an Open instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.path),
		[]byte(app.permission),
	})

	if err != nil {
		return nil, err
	}

	return createOpen(*pHash, app.path, app.permission), nil
}
