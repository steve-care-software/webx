package trees

import "errors"

type builder struct {
	list []Tree
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
func (app *builder) WithList(list []Tree) Builder {
	app.list = list
	return app
}

// Now builds a new Trees instance
func (app *builder) Now() (Trees, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Tree in order to build a Trees instance")
	}

	return createTrees(app.list), nil
}
