package containers

import "errors"

type builder struct {
	list []Container
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
func (app *builder) WithList(list []Container) Builder {
	app.list = list
	return app
}

// Now builds a new Containers instance
func (app *builder) Now() (Containers, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Container in order to build a Containers instance")
	}

	mp := map[string]Container{}
	for _, oneContainer := range app.list {
		keyname := oneContainer.Keyname()
		mp[keyname] = oneContainer
	}

	return createContainers(
		mp,
		app.list,
	), nil
}
