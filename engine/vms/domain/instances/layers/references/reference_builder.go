package references

import (
	"errors"
	"path/filepath"

	"github.com/steve-care-software/webx/engine/states/domain/hash"
)

type referenceBuilder struct {
	hashAdapter hash.Adapter
	variable    string
	path        []string
}

func createReferenceBuilder(
	hashAdapter hash.Adapter,
) ReferenceBuilder {
	out := referenceBuilder{
		hashAdapter: hashAdapter,
		variable:    "",
		path:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *referenceBuilder) Create() ReferenceBuilder {
	return createReferenceBuilder(
		app.hashAdapter,
	)
}

// WithVariable adds a variable to the builder
func (app *referenceBuilder) WithVariable(variable string) ReferenceBuilder {
	app.variable = variable
	return app
}

// WithPath adds a path to the builder
func (app *referenceBuilder) WithPath(path []string) ReferenceBuilder {
	app.path = path
	return app
}

// Now builds a new Reference instance
func (app *referenceBuilder) Now() (Reference, error) {
	if app.path != nil && len(app.path) <= 0 {
		app.path = nil
	}

	if app.path == nil {
		return nil, errors.New("the path is mandatory in order to build a Reference instance")
	}

	if app.variable == "" {
		return nil, errors.New("the variable is mandatory in order to build a Reference instance")
	}

	path := filepath.Join(app.path...)
	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.variable),
		[]byte(path),
	})

	if err != nil {
		return nil, err
	}

	return createReference(*pHash, app.variable, app.path), nil
}
