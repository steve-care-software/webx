package tokens

import (
	"errors"

	"github.com/steve-care-software/webx/engine/domain/programs/grammars/constants/tokens/elements"
)

type tokenBuilder struct {
	name    string
	element elements.Element
	amount  uint
}

func createTokenBuilder() TokenBuilder {
	out := tokenBuilder{
		name:    "",
		element: nil,
		amount:  0,
	}

	return &out
}

// Create initializes the builder
func (app *tokenBuilder) Create() TokenBuilder {
	return createTokenBuilder()
}

// WithName adds a name to the builder
func (app *tokenBuilder) WithName(name string) TokenBuilder {
	app.name = name
	return app
}

// WithElement adds an element to the builder
func (app *tokenBuilder) WithElement(element elements.Element) TokenBuilder {
	app.element = element
	return app
}

// WithAmount adds an amount to the builder
func (app *tokenBuilder) WithAmount(amount uint) TokenBuilder {
	app.amount = amount
	return app
}

// Now builds a new Token instance
func (app *tokenBuilder) Now() (Token, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Token instance")
	}

	if app.element == nil {
		return nil, errors.New("the element is mandatory in order to build a Token instance")
	}

	if app.amount <= 0 {
		return nil, errors.New("the amount is mandatory in order to build a Token instance")
	}

	return createToken(
		app.name,
		app.element,
		app.amount,
	), nil
}
