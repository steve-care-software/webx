package singles

import (
	"errors"
)

type builder struct {
	list []Single
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
func (app *builder) WithList(list []Single) Builder {
	app.list = list
	return app
}

// Now builds a new Singles instance
func (app *builder) Now() (Singles, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Single in order to build a Singles instance")
	}

	return createSingles(app.list), nil
}
