package parameters

import (
	"errors"
)

type parameterBuilder struct {
	token  string
	pIndex *uint
	name   string
}

func createParameterBuilder() ParameterBuilder {
	out := parameterBuilder{
		token:  "",
		pIndex: nil,
		name:   "",
	}

	return &out
}

// Create initializes the parameterBuilder
func (app *parameterBuilder) Create() ParameterBuilder {
	return createParameterBuilder()
}

// WithToken adds a token to the parameterBuilder
func (app *parameterBuilder) WithToken(token string) ParameterBuilder {
	app.token = token
	return app
}

// WithIndex adds an index to the parameterBuilder
func (app *parameterBuilder) WithIndex(index uint) ParameterBuilder {
	app.pIndex = &index
	return app
}

// WithName adds a name to the parameterBuilder
func (app *parameterBuilder) WithName(name string) ParameterBuilder {
	app.name = name
	return app
}

// Now builds a new Parameter instance
func (app *parameterBuilder) Now() (Parameter, error) {
	if app.token == "" {
		return nil, errors.New("the token is mandatory in order to build a Parameter instance")
	}

	if app.pIndex == nil {
		return nil, errors.New("the index is mandatory in order to build a Parameter instance")
	}

	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Parameter instance")
	}

	return createParameter(
		app.token,
		*app.pIndex,
		app.name,
	), nil
}
