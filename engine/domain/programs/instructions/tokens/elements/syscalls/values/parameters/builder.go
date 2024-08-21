package parameters

import (
	"errors"
)

type builder struct {
	element string
	pIndex  *uint
	name    string
}

func createBuilder() Builder {
	out := builder{
		element: "",
		pIndex:  nil,
		name:    "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithElement adds a element to the builder
func (app *builder) WithElement(element string) Builder {
	app.element = element
	return app
}

// WithIndex adds an index to the builder
func (app *builder) WithIndex(index uint) Builder {
	app.pIndex = &index
	return app
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// Now builds a new Parameter instance
func (app *builder) Now() (Parameter, error) {
	if app.element == "" {
		return nil, errors.New("the element is mandatory in order to build a Parameter instance")
	}

	if app.pIndex == nil {
		return nil, errors.New("the index is mandatory in order to build a Parameter instance")
	}

	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Parameter instance")
	}

	return createParameter(
		app.element,
		*app.pIndex,
		app.name,
	), nil
}
