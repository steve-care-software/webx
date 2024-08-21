package values

import (
	"errors"

	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/executions/parameters"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens"
)

type valueBuilder struct {
	parameter parameters.Parameter
	token     tokens.Token
}

func createValueBuilder() ValueBuilder {
	out := valueBuilder{
		parameter: nil,
		token:     nil,
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
func (app *valueBuilder) WithToken(token tokens.Token) ValueBuilder {
	app.token = token
	return app
}

// Now builds a new Value instance
func (app *valueBuilder) Now() (Value, error) {
	if app.parameter != nil {
		return createValueWithParameter(app.parameter), nil
	}

	if app.token != nil {
		return createValueWithToken(app.token), nil
	}

	return nil, errors.New("the Value is invalid")
}
