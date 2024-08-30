package validations

import "errors"

type builder struct {
	list []Validation
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
func (app *builder) WithList(list []Validation) Builder {
	app.list = list
	return app
}

// Now builds a new Validations instance
func (app *builder) Now() (Validations, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Validation in order to build a Validations instance")
	}

	return createValidations(
		app.list,
	), nil
}
