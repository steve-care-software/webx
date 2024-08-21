package values

import "errors"

type builder struct {
	list []Value
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
func (app *builder) WithList(list []Value) Builder {
	app.list = list
	return app
}

// Now builds a new Values instance
func (app *builder) Now() (Values, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Value in order to build a Values instance")
	}

	return createValues(app.list), nil
}
