package actions

import (
	"errors"
	"path/filepath"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications"
)

type actionBuilder struct {
	hashAdapter   hash.Adapter
	path          []string
	modifications modifications.Modifications
}

func createActionBuilder(
	hashAdapter hash.Adapter,
) ActionBuilder {
	out := actionBuilder{
		hashAdapter:   hashAdapter,
		path:          nil,
		modifications: nil,
	}

	return &out
}

// Create initializes the builder
func (app *actionBuilder) Create() ActionBuilder {
	return createActionBuilder(
		app.hashAdapter,
	)
}

// WithPath adds a path to the builder
func (app *actionBuilder) WithPath(path []string) ActionBuilder {
	app.path = path
	return app
}

// WithModifications adds modifications to the builder
func (app *actionBuilder) WithModifications(modifications modifications.Modifications) ActionBuilder {
	app.modifications = modifications
	return app
}

// Now builds a new Action instance
func (app *actionBuilder) Now() (Action, error) {
	if app.path != nil && len(app.path) <= 0 {
		app.path = nil
	}

	if app.path == nil {
		return nil, errors.New("the path is mandatory in order to build an Action instance")
	}

	if app.modifications == nil {
		return nil, errors.New("the modifications is mandatory in order to build an Action instance")
	}

	path := filepath.Join(app.path...)
	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(path),
		app.modifications.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createAction(*pHash, app.path, app.modifications), nil
}
