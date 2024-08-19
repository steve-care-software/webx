package suites

import (
	"errors"
)

type builder struct {
	list []Suite
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
func (app *builder) WithList(list []Suite) Builder {
	app.list = list
	return app
}

// Now builds a new Suites instance
func (app *builder) Now() (Suites, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Suite in order to build a Suites instance")
	}

	return createSuites(app.list), nil
}
