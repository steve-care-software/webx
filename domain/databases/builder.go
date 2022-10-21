package databases

import (
	"errors"
)

type builder struct {
	list []Database
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
func (app *builder) WithList(list []Database) Builder {
	app.list = list
	return app
}

// Now builds a new Databases instance
func (app *builder) Now() (Databases, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Database in order to build an Databases instance")
	}

	return createDatabases(app.list), nil
}
