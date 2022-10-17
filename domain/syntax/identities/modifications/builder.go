package modifications

import (
	"errors"
)

type builder struct {
	list []Modification
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
func (app *builder) WithList(list []Modification) Builder {
	app.list = list
	return app
}

// Now builds a new Modifications instance
func (app *builder) Now() (Modifications, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Modification in order to build an Modifications instance")
	}

	return createModifications(app.list), nil
}
