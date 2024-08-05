package updates

import "errors"

type builder struct {
	name        string
	description string
}

func createBuilder() Builder {
	out := builder{
		name:        "",
		description: "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithDescription adds a description to the builder
func (app *builder) WithDescription(description string) Builder {
	app.description = description
	return app
}

// Now builds a new Update instance
func (app *builder) Now() (Update, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build an Update instance")
	}

	return createUpdate(
		app.name,
		app.description,
	), nil
}
