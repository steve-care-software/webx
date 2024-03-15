package connections

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type fieldBuilder struct {
	hashAdapter hash.Adapter
	name        string
	path        []string
}

func createFieldBuilder(
	hashAdapter hash.Adapter,
) FieldBuilder {
	out := fieldBuilder{
		hashAdapter: hashAdapter,
		name:        "",
		path:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *fieldBuilder) Create() FieldBuilder {
	return createFieldBuilder(
		app.hashAdapter,
	)
}

// WithName adds a name to the builder
func (app *fieldBuilder) WithName(name string) FieldBuilder {
	app.name = name
	return app
}

// WithPath adds a path to the builder
func (app *fieldBuilder) WithPath(path []string) FieldBuilder {
	app.path = path
	return app
}

// Now builds a new Field instance
func (app *fieldBuilder) Now() (Field, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Field instance")
	}

	if app.path != nil && len(app.path) <= 0 {
		app.path = nil
	}

	if app.path == nil {
		return nil, errors.New("the path is mandatory in order to build a Field instance")
	}

	data := [][]byte{
		[]byte(app.name),
	}

	for _, onePath := range app.path {
		data = append(data, []byte(onePath))
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createField(*pHash, app.name, app.path), nil
}
