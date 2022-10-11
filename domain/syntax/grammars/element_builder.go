package grammars

import (
	"errors"

	"github.com/steve-care-software/syntax/domain/syntax/grammars/cardinalities"
	"github.com/steve-care-software/syntax/domain/syntax/grammars/values"
)

type elementBuilder struct {
	cardinality cardinalities.Cardinality
	value       values.Value
	external    External
	instance    Instance
}

func createElementBuilder() ElementBuilder {
	out := elementBuilder{
		cardinality: nil,
		value:       nil,
		external:    nil,
		instance:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *elementBuilder) Create() ElementBuilder {
	return createElementBuilder()
}

// WithCardinality adds a cardinality to the builder
func (app *elementBuilder) WithCardinality(cardinality cardinalities.Cardinality) ElementBuilder {
	app.cardinality = cardinality
	return app
}

// WithValue adds a value to the builder
func (app *elementBuilder) WithValue(value values.Value) ElementBuilder {
	app.value = value
	return app
}

// WithExternal adds an external grammar to the builder
func (app *elementBuilder) WithExternal(external External) ElementBuilder {
	app.external = external
	return app
}

// WithInstance adds an instance to the builder
func (app *elementBuilder) WithInstance(instance Instance) ElementBuilder {
	app.instance = instance
	return app
}

// Now builds a new Element instance
func (app *elementBuilder) Now() (Element, error) {
	if app.value != nil {
		content := createElementContentWithValue(app.value)
		return createElement(content, app.cardinality), nil
	}

	if app.external != nil {
		content := createElementContentWithExternalToken(app.external)
		return createElement(content, app.cardinality), nil
	}

	if app.instance != nil {
		content := createElementContentWithInstance(app.instance)
		return createElement(content, app.cardinality), nil
	}

	return nil, errors.New("the Element is invalid")
}
