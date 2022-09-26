package compilers

import (
	"errors"

	"github.com/steve-care-software/syntax/domain/syntax/bytes/grammars"
)

type elementBuilder struct {
	grammar     grammars.Grammar
	composition Composition
}

func createElementBuilder() ElementBuilder {
	out := elementBuilder{
		grammar:     nil,
		composition: nil,
	}

	return &out
}

// Create initializes the builder
func (app *elementBuilder) Create() ElementBuilder {
	return createElementBuilder()
}

// WithGrammar adds a grammar to the builder
func (app *elementBuilder) WithGrammar(grammar grammars.Grammar) ElementBuilder {
	app.grammar = grammar
	return app
}

// WithComposition adds a composition to the builder
func (app *elementBuilder) WithComposition(composition Composition) ElementBuilder {
	app.composition = composition
	return app
}

// Now builds a new Element instance
func (app *elementBuilder) Now() (Element, error) {
	if app.grammar == nil {
		return nil, errors.New("the grammar is mandatory in order to build an Element instance")
	}

	if app.composition == nil {
		return nil, errors.New("the composition is mandatory in order to build an Element instance")
	}

	return createElement(app.grammar, app.composition), nil
}
