package applications

import "errors"

type builder struct {
	module string
	name   string
}

func createBuilder() Builder {
	out := builder{
		module: "",
		name:   "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithModule adds a module to the builder
func (app *builder) WithModule(module string) Builder {
	app.module = module
	return app
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// Now builds a new Application instance
func (app *builder) Now() (Application, error) {
	if app.module == "" {
		return nil, errors.New("the module is mandatory in order to build an Application instance")
	}

	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build an Application instance")
	}

	return createApplication(app.module, app.name), nil
}
