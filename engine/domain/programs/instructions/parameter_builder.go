package instructions

import (
	"errors"
)

type parameterBuilder struct {
	name  string
	value Value
}

func createParameterBuilder() ParameterBuilder {
	out := parameterBuilder{
		name:  "",
		value: nil,
	}

	return &out
}

// Create initializes the builder
func (app *parameterBuilder) Create() ParameterBuilder {
	return createParameterBuilder()
}

// WithName adds a name to the builder
func (app *parameterBuilder) WithName(name string) ParameterBuilder {
	app.name = name
	return app
}

// WithValue adds a value to the builder
func (app *parameterBuilder) WithValue(value Value) ParameterBuilder {
	app.value = value
	return app
}

// Now builds a new Parameter instance
func (app *parameterBuilder) Now() (Parameter, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Parameter instance")
	}

	if app.value == nil {
		return nil, errors.New("the value is mandatory in order to build a Parameter instance")
	}

	return createParameter(
		app.name,
		app.value,
	), nil
}
