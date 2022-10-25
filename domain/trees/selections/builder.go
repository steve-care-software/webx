package selections

import "errors"

type builder struct {
	treeName string
	list     []Selection
}

func createBuilder() Builder {
	out := builder{
		treeName: "",
		list:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithTreeName adds a treeName to the builder
func (app *builder) WithTreeName(treeName string) Builder {
	app.treeName = treeName
	return app
}

// WithList adds a selection list to the builder
func (app *builder) WithList(list []Selection) Builder {
	app.list = list
	return app
}

// Now builds a new Selections instance
func (app *builder) Now() (Selections, error) {
	if app.treeName == "" {
		return nil, errors.New("the treeName is mandatory in order to build a Selections instance")
	}

	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Selection in order to build a Selections instance")
	}

	return createSelections(app.treeName, app.list), nil
}
