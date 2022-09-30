package uncovers

import "errors"

type builder struct {
	list []Uncover
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
func (app *builder) WithList(list []Uncover) Builder {
	app.list = list
	return app
}

// Now builds a new Uncovers instance
func (app *builder) Now() (Uncovers, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Uncover in order to build a Uncovers instance")
	}

	return createUncovers(app.list), nil
}
