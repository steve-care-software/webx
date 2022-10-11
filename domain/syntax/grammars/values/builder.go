package values

import "errors"

type builder struct {
	name    string
	pNumber *byte
}

func createBuilder() Builder {
	out := builder{
		name:    "",
		pNumber: nil,
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

// WithNumber adds a number to the builder
func (app *builder) WithNumber(number byte) Builder {
	app.pNumber = &number
	return app
}

// Now builds a new Value instance
func (app *builder) Now() (Value, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Value instance")
	}

	if app.pNumber == nil {
		return nil, errors.New("the value is mandatory in order to build a Value instance")
	}

	return createValue(app.name, *app.pNumber), nil
}
