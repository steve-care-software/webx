package modules

import "errors"

type builder struct {
	name string
	fn   ExecuteFn
}

func createBuilder() Builder {
	out := builder{
		name: "",
		fn:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithFunc adds a func to the builder
func (app *builder) WithFunc(fn ExecuteFn) Builder {
	app.fn = fn
	return app
}

// Now builds a new Module instance
func (app *builder) Now() (Module, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Module instance")
	}

	if app.fn == nil {
		return nil, errors.New("the execute func is mandatory in order to build a Module instance")
	}

	return createModule(app.name, app.fn), nil
}
