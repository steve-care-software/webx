package trees

import (
	"errors"

	"github.com/steve-care-software/syntax/domain/syntax/bytes/grammars"
)

type elementBuilder struct {
	grammar  grammars.Element
	contents Contents
}

func createElementBuilder() ElementBuilder {
	out := elementBuilder{
		grammar:  nil,
		contents: nil,
	}

	return &out
}

// Create initializes the builder
func (app *elementBuilder) Create() ElementBuilder {
	return createElementBuilder()
}

// WithGrammar adds a grammar to the builder
func (app *elementBuilder) WithGrammar(grammar grammars.Element) ElementBuilder {
	app.grammar = grammar
	return app
}

// WithContents adds a contents to the builder
func (app *elementBuilder) WithContents(contents Contents) ElementBuilder {
	app.contents = contents
	return app
}

// Now builds a new Element instance
func (app *elementBuilder) Now() (Element, error) {
	if app.grammar == nil {
		return nil, errors.New("the grammar is mandatory in order to build an Element instance")
	}

	if app.contents == nil {
		return nil, errors.New("the contents is mandatory in order to build an Element instance")
	}

	return createElement(app.grammar, app.contents), nil
}
