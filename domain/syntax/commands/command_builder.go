package commands

import (
	"errors"

	"github.com/steve-care-software/syntax/domain/syntax/bytes/grammars"
)

type commandBuilder struct {
	grammar grammars.Grammar
	content Content
}

func createCommandBuilder() CommandBuilder {
	out := commandBuilder{
		grammar: nil,
		content: nil,
	}

	return &out
}

// Create initializes the builder
func (app *commandBuilder) Create() CommandBuilder {
	return createCommandBuilder()
}

// WithGrammar adds a grammar to the builder
func (app *commandBuilder) WithGrammar(grammar grammars.Grammar) CommandBuilder {
	app.grammar = grammar
	return app
}

// WithContent adds content to the builder
func (app *commandBuilder) WithContent(content Content) CommandBuilder {
	app.content = content
	return app
}

// Now builds a new Command instance
func (app *commandBuilder) Now() (Command, error) {
	if app.grammar == nil {
		return nil, errors.New("the grammar is mandatory in order to build a Command instance")
	}

	if app.content == nil {
		return nil, errors.New("the content is mandatory in order to build a Command instance")
	}

	return createCommand(app.grammar, app.content), nil
}
