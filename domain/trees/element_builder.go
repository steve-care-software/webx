package trees

import (
	"errors"

	"github.com/steve-care-software/webx/domain/grammars"
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
	if app.contents == nil {
		return nil, errors.New("the contents is mandatory in order to build an Element instance")
	}

	if app.grammar != nil {
		return createElementWithGrammar(app.contents, app.grammar), nil
	}

	return createElement(app.contents), nil
}
