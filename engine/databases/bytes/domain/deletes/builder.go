package deletes

import "errors"

type builder struct {
	list []Delete
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
func (app *builder) WithList(list []Delete) Builder {
	app.list = list
	return app
}

// Now builds a new Deletes instance
func (app *builder) Now() (Deletes, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Delete in order to build a Deletes instance")
	}

	return createDeletes(
		app.list,
	), nil
}
