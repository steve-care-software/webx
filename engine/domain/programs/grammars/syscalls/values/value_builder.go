package values

import (
	"errors"

	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/executions/parameters"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens/elements"
)

type valueBuilder struct {
	parameter parameters.Parameter
	element   elements.Element
}

func createValueBuilder() ValueBuilder {
	out := valueBuilder{
		parameter: nil,
		element:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *valueBuilder) Create() ValueBuilder {
	return createValueBuilder()
}

// WithParameter adds a parameter to the builder
func (app *valueBuilder) WithParameter(parameter parameters.Parameter) ValueBuilder {
	app.parameter = parameter
	return app
}

// WithElement adds an element to the builder
func (app *valueBuilder) WithElement(element elements.Element) ValueBuilder {
	app.element = element
	return app
}

// Now builds a new Value instance
func (app *valueBuilder) Now() (Value, error) {
	if app.parameter != nil {
		return createValueWithParameter(app.parameter), nil
	}

	if app.element != nil {
		return createValueWithElement(app.element), nil
	}

	return nil, errors.New("the Value is invalid")
}
