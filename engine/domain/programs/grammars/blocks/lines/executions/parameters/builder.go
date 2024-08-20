package parameters

import (
	"errors"
)

type builder struct {
	list []Parameter
}

func createBuilder() Builder {
	out := builder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithList adds a list to the builder
func (app *builder) WithList(list []Parameter) Builder {
	app.list = list
	return app
}

// Now builds a new Parameters instance
func (app *builder) Now() (Parameters, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Parameter in order to build a Parameters instance")
	}

	return createParameters(app.list), nil
}
