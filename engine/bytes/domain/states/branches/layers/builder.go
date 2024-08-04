package layers

import "errors"

type builder struct {
	list []Layer
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
func (app *builder) WithList(list []Layer) Builder {
	app.list = list
	return app
}

// Now builds a new Layers instance
func (app *builder) Now() (Layers, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Layer in order to build a Layers instance")
	}

	return createLayers(
		app.list,
	), nil
}
