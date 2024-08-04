package delimiters

import "errors"

type builder struct {
	list []Delimiter
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
func (app *builder) WithList(list []Delimiter) Builder {
	app.list = list
	return app
}

// Now builds a new Delimiters instance
func (app *builder) Now() (Delimiters, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Delimiter in order to build an Delimiters instance")
	}

	return createDelimiters(
		app.list,
	), nil
}
