package modules

import "errors"

type builder struct {
	pIndex *uint
	name   []byte
}

func createBuilder() Builder {
	out := builder{
		pIndex: nil,
		name:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithIndex adds an index to the builder
func (app *builder) WithIndex(index uint) Builder {
	app.pIndex = &index
	return app
}

// WithName adds a name to the builder
func (app *builder) WithName(name []byte) Builder {
	app.name = name
	return app
}

// Now builds a new Module instance
func (app *builder) Now() (Module, error) {
	if app.pIndex == nil {
		return nil, errors.New("the index is mandatory in order to build a Module instance")
	}

	if app.name == nil {
		return nil, errors.New("the name is mandatory in order to build a Module instance")
	}

	return createModule(*app.pIndex, app.name), nil
}
