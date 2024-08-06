package originals

import "errors"

type builder struct {
	name        string
	description string
	isDeleted   bool
}

func createBuilder() Builder {
	out := builder{
		name:        "",
		description: "",
		isDeleted:   false,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// IsDeleted flags the builder as deleted
func (app *builder) IsDeleted() Builder {
	app.isDeleted = true
	return app
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

// Now builds a new Original instance
func (app *builder) Now() (Original, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build an Original instance")
	}

	return createOriginal(
		app.name,
		app.description,
		app.isDeleted,
	), nil
}
