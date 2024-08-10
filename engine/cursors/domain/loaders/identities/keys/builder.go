package keys

import (
	"errors"
)

type builder struct {
	list []Key
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
func (app *builder) WithList(list []Key) Builder {
	app.list = list
	return app
}

// Now builds a new Keys instance
func (app *builder) Now() (Keys, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Key in order to build a Keys instance")
	}

	return createKeys(app.list), nil
}
