package storages

import "errors"

type builder struct {
	list []Storage
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
func (app *builder) WithList(list []Storage) Builder {
	app.list = list
	return app
}

// Now builds a new Storages instance
func (app *builder) Now() (Storages, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Storage in order to build a Storages instance")
	}

	return createStorages(
		app.list,
	), nil
}
