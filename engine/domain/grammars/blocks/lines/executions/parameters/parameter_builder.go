package parameters

import (
	"errors"

	"github.com/steve-care-software/webx/engine/domain/grammars/blocks/lines/tokens/elements"
)

type parameterBuilder struct {
	element elements.Element
	pIndex  *uint
	name    string
}

func createParameterBuilder() ParameterBuilder {
	out := parameterBuilder{
		element: nil,
		pIndex:  nil,
		name:    "",
	}

	return &out
}

// Create initializes the builder
func (app *parameterBuilder) Create() ParameterBuilder {
	return createParameterBuilder()
}

// WithElement adds an element to the builder
func (app *parameterBuilder) WithElement(element elements.Element) ParameterBuilder {
	app.element = element
	return app
}

// WithIndex adds an index to the builder
func (app *parameterBuilder) WithIndex(index uint) ParameterBuilder {
	app.pIndex = &index
	return app
}

// WithName adds a name to the builder
func (app *parameterBuilder) WithName(name string) ParameterBuilder {
	app.name = name
	return app
}

// Now builds a new Parameter instance
func (app *parameterBuilder) Now() (Parameter, error) {
	if app.element == nil {
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
