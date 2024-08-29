package variables

import (
	"errors"
)

type builder struct {
	list []Variable
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
func (app *builder) WithList(list []Variable) Builder {
	app.list = list
	return app
}

// Now builds a new Variables instance
func (app *builder) Now() (Variables, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Variable in order to build a Variables instance")
	}

	mp := map[string]Variable{}
	for _, oneVariable := range app.list {
		keyname := oneVariable.Name()
		mp[keyname] = oneVariable
	}

	return createVariables(
		app.list,
		mp,
	), nil
}
