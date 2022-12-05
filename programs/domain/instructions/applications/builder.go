package applications

import "errors"

type builder struct {
	module []byte
	name   []byte
}

func createBuilder() Builder {
	out := builder{
		module: nil,
		name:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithModule adds a module to the builder
func (app *builder) WithModule(module []byte) Builder {
	app.module = module
	return app
}

// WithName adds a name to the builder
func (app *builder) WithName(name []byte) Builder {
	app.name = name
	return app
}

// Now builds a new Application instance
func (app *builder) Now() (Application, error) {
	if app.module == nil {
		return nil, errors.New("the module is mandatory in order to build an Application instance")
	}

	if app.name == nil {
		return nil, errors.New("the name is mandatory in order to build an Application instance")
	}

	return createApplication(app.module, app.name), nil
}
