package tokens

import (
	"errors"

	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens/cardinalities"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens/elements"
)

type tokenBuilder struct {
	element     elements.Element
	cardinality cardinalities.Cardinality
	isReverse   bool
}

func createTokenBuilder() TokenBuilder {
	out := tokenBuilder{
		element:     nil,
		cardinality: nil,
		isReverse:   false,
	}

	return &out
}

// Create initializes the builder
func (app *tokenBuilder) Create() TokenBuilder {
	return createTokenBuilder()
}

// WithElement adds an element to the builder
func (app *tokenBuilder) WithElement(element elements.Element) TokenBuilder {
	app.element = element
	return app
}

// WithCardinality adds a cardinality to the builder
func (app *tokenBuilder) WithCardinality(cardinality cardinalities.Cardinality) TokenBuilder {
	app.cardinality = cardinality
	return app
}

// IsReverse flags the builder as isReverse
func (app *tokenBuilder) IsReverse() TokenBuilder {
	app.isReverse = true
	return app
}

// Now builds a new Token instance
func (app *tokenBuilder) Now() (Token, error) {
	if app.element == nil {
		return nil, errors.New("the element is mandatory in order to build a Token instance")
	}

	if app.cardinality == nil {
		return nil, errors.New("the cardinality is mandatory in order to build a Token instance")
	}

	return createToken(
		app.element,
		app.cardinality,
		app.isReverse,
	), nil
}
