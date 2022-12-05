package modules

import "errors"

type moduleBuilder struct {
	name []byte
	fn   ExecuteFn
}

func createModuleBuilder() ModuleBuilder {
	out := moduleBuilder{
		name: nil,
		fn:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *moduleBuilder) Create() ModuleBuilder {
	return createModuleBuilder()
}

// WithName adds a name to the builder
func (app *moduleBuilder) WithName(name []byte) ModuleBuilder {
	app.name = name
	return app
}

// WithFunc adds a func to the builder
func (app *moduleBuilder) WithFunc(fn ExecuteFn) ModuleBuilder {
	app.fn = fn
	return app
}

// Now builds a new Module instance
func (app *moduleBuilder) Now() (Module, error) {
	if app.name == nil {
		return nil, errors.New("the name is mandatory in order to build a Module instance")
	}

	if app.fn == nil {
		return nil, errors.New("the execute func is mandatory in order to build a Module instance")
	}

	return createModule(app.name, app.fn), nil
}
