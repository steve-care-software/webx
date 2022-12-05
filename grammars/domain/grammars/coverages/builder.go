package coverages

import "errors"

type builder struct {
	list []Coverage
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
func (app *builder) WithList(list []Coverage) Builder {
	app.list = list
	return app
}

// Now builds a new Coverages instance
func (app *builder) Now() (Coverages, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Coverage in order to build a Coverages instance")
	}

	return createCoverages(app.list), nil
}
