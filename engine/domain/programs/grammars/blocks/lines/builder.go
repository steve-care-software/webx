package lines

import (
	"errors"
)

type builder struct {
	list []Line
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
func (app *builder) WithList(list []Line) Builder {
	app.list = list
	return app
}

// Now builds a new Lines instance
func (app *builder) Now() (Lines, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Line in order to build a Lines instance")
	}

	return createLines(app.list), nil
}
