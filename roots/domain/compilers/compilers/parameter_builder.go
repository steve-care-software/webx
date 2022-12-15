package compilers

import (
	"errors"

	"github.com/steve-care-software/webx/roots/domain/selectors/selectors"
)

type parameterBuilder struct {
	pIndex   *uint
	selector selectors.Selector
}

func createParameterBuilder() ParameterBuilder {
	out := parameterBuilder{
		pIndex:   nil,
		selector: nil,
	}

	return &out
}

// Create initializes the builder
func (app *parameterBuilder) Create() ParameterBuilder {
	return createParameterBuilder()
}

// WithIndex adds an index to the builder
func (app *parameterBuilder) WithIndex(index uint) ParameterBuilder {
	app.pIndex = &index
	return app
}

// WithSelector adds a selector to the builder
func (app *parameterBuilder) WithSelector(selector selectors.Selector) ParameterBuilder {
	app.selector = selector
	return app
}

// Now builds a new Parameter instance
func (app *parameterBuilder) Now() (Parameter, error) {
	if app.pIndex == nil {
		return nil, errors.New("the index is mandatory in order to build a Parameter instance")
	}

	if app.selector == nil {
		return nil, errors.New("the selector is mandatory in order to build a Parameter instance")
	}

	return createPrameter(*app.pIndex, app.selector), nil
}
