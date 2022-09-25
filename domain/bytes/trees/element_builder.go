package trees

import (
	"errors"

	"github.com/steve-care-software/syntax/domain/bytes/grammars"
)

type elementBuilder struct {
	grammar grammars.Element
	content Content
	amount  uint
}

func createElementBuilder() ElementBuilder {
	out := elementBuilder{
		grammar: nil,
		content: nil,
		amount:  0,
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

// WithContent adds a content to the builder
func (app *elementBuilder) WithContent(content Content) ElementBuilder {
	app.content = content
	return app
}

// WithAmount adds an amount to the builder
func (app *elementBuilder) WithAmount(amount uint) ElementBuilder {
	app.amount = amount
	return app
}

// Now builds a new Element instance
func (app *elementBuilder) Now() (Element, error) {
	if app.grammar == nil {
		return nil, errors.New("the grammar is mandatory in order to build an Element instance")
	}

	if app.content == nil {
		return nil, errors.New("the content is mandatory in order to build an Element instance")
	}

	if app.amount <= 0 {
		return nil, errors.New("the amount must be greater than zero (0) in order to build an Element instance")
	}

	return createElement(app.grammar, app.content, app.amount), nil
}
