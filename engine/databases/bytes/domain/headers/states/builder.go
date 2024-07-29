package states

import "errors"

type builder struct {
	list []State
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
func (app *builder) WithList(list []State) Builder {
	app.list = list
	return app
}

// Now builds a new States instance
func (app *builder) Now() (States, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 State in order to build a States instance")
	}

	return createStates(
		app.list,
	), nil
}
