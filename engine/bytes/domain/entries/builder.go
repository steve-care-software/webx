package entries

import "errors"

type builder struct {
	list []Entry
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
func (app *builder) WithList(list []Entry) Builder {
	app.list = list
	return app
}

// Now builds a new Entries instance
func (app *builder) Now() (Entries, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Entry in order to build an Entries instance")
	}

	return createEntries(
		app.list,
	), nil
}
