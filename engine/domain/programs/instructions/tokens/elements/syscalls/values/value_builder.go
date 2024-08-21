package values

import (
	"errors"

	"github.com/steve-care-software/webx/engine/domain/programs/instructions/tokens/elements/syscalls/values/parameters"
)

type valueBuilder struct {
	parameter parameters.Parameter
	token     string
}

func createValueBuilder() ValueBuilder {
	out := valueBuilder{
		parameter: nil,
		token:     "",
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

// WithToken adds an token to the builder
func (app *valueBuilder) WithToken(token string) ValueBuilder {
	app.token = token
	return app
}

// Now builds a new Value instance
func (app *valueBuilder) Now() (Value, error) {
	if app.parameter != nil {
		return createValueWithParameter(app.parameter), nil
	}

	if app.token != "" {
		return createValueWithToken(app.token), nil
	}

	return nil, errors.New("the Value is invalid")
}
