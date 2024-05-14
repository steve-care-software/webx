package actions

import (
	"errors"
	"path/filepath"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type builder struct {
	hashAdapter   hash.Adapter
	path          string
	modifications string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter:   hashAdapter,
		path:          "",
		modifications: "",
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

// WithModifications adds a modifications to the builder
func (app *builder) WithModifications(modifications string) Builder {
	app.modifications = modifications
	return app
}

// Now builds a new Action instance
func (app *builder) Now() (Action, error) {
	if app.path == "" {
		return nil, errors.New("the path is mandatory in order to build an Action instance")
	}

	if app.modifications == "" {
		return nil, errors.New("the modifications is mandatory in order to build an Action instance")
	}

	path := filepath.Join(app.path)
	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(path),
		[]byte(app.modifications),
	})

	if err != nil {
		return nil, err
	}

	return createAction(*pHash, app.path, app.modifications), nil
}
