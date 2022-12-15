package modules

import "errors"

type moduleBuilder struct {
	pIndex *uint
	fn     ExecuteFn
}

func createModuleBuilder() ModuleBuilder {
	out := moduleBuilder{
		pIndex: nil,
		fn:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *moduleBuilder) Create() ModuleBuilder {
	return createModuleBuilder()
}

// WithIndex adds an index the builder
func (app *moduleBuilder) WithIndex(index uint) ModuleBuilder {
	app.pIndex = &index
	return app
}

// WithFunc adds a func to the builder
func (app *moduleBuilder) WithFunc(fn ExecuteFn) ModuleBuilder {
	app.fn = fn
	return app
}

// Now builds a new Module instance
func (app *moduleBuilder) Now() (Module, error) {
	if app.pIndex == nil {
		return nil, errors.New("the index is mandatory in order to build a Module instance")
	}

	if app.fn == nil {
		return nil, errors.New("the execute func is mandatory in order to build a Module instance")
	}

	return createModule(*app.pIndex, app.fn), nil
}
