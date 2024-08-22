package instructions

import (
	"errors"
)

type tokenBuilder struct {
	name     string
	elements Elements
}

func createTokenBuilder() TokenBuilder {
	out := tokenBuilder{
		name:     "",
		elements: nil,
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

// WithElements add elements to the builder
func (app *tokenBuilder) WithElements(elements Elements) TokenBuilder {
	app.elements = elements
	return app
}

// Now builds a new Token instance
func (app *tokenBuilder) Now() (Token, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Token instance")
	}

	if app.elements == nil {
		return nil, errors.New("the elements is mandatory in order to build a Token instance")
	}

	return createToken(app.name, app.elements), nil
}
