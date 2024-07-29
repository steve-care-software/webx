package retrievals

import "errors"

type builder struct {
	list []Retrieval
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
func (app *builder) WithList(list []Retrieval) Builder {
	app.list = list
	return app
}

// Now builds a new Retrievals instance
func (app *builder) Now() (Retrievals, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Retrieval in order to build an Retrievals instance")
	}

	return createRetrievals(
		app.list,
	), nil
}
