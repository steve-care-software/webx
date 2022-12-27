package modules

import (
	"errors"
)

type builder struct {
	list []Module
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
func (app *builder) WithList(list []Module) Builder {
	app.list = list
	return app
}

// Now builds a new Modules instance
func (app *builder) Now() (Modules, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Module in order to build a Modules instance")
	}

	mp := map[uint]Module{}
	for _, oneModule := range app.list {
		index := oneModule.Index()
		mp[index] = oneModule
	}

	return createModules(app.list, mp), nil
}
