package connections

import "errors"

type fieldBuilder struct {
	name string
	path []string
}

func createFieldBuilder() FieldBuilder {
	out := fieldBuilder{
		name: "",
		path: nil,
	}

	return &out
}

// Create initializes the builder
func (app *fieldBuilder) Create() FieldBuilder {
	return createFieldBuilder()
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

	return createField(app.name, app.path), nil
}
