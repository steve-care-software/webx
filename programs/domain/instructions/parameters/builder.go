package parameters

import "errors"

type builder struct {
	name    []byte
	isInput bool
}

func createBuilder() Builder {
	out := builder{
		name:    nil,
		isInput: false,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithName adds a name to the builder
func (app *builder) WithName(name []byte) Builder {
	app.name = name
	return app
}

// IsInput flags the builder as an input
func (app *builder) IsInput() Builder {
	app.isInput = true
	return app
}

// Now builds a new Parameter instance
func (app *builder) Now() (Parameter, error) {
	if app.name == nil {
		return nil, errors.New("the name is mandatory in order to build a Parameter instance")
	}

	if app.isInput {
		return createParameterWithInput(app.name), nil
	}

	return createParameterWithOutput(app.name), nil
}
