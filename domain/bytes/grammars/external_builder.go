package grammars

import "errors"

type externalBuilder struct {
	name    string
	grammar Grammar
}

func createExternalBuilder() ExternalBuilder {
	out := externalBuilder{
		name:    "",
		grammar: nil,
	}

	return &out
}

// Create initializes the builder
func (app *externalBuilder) Create() ExternalBuilder {
	return createExternalBuilder()
}

// WithName adds a name to the builder
func (app *externalBuilder) WithName(name string) ExternalBuilder {
	app.name = name
	return app
}

// WithGrammar adds a grammar to the builder
func (app *externalBuilder) WithGrammar(grammar Grammar) ExternalBuilder {
	app.grammar = grammar
	return app
}

// Now builds a new External instance
func (app *externalBuilder) Now() (External, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build an External instance")
	}

	if app.grammar == nil {
		return nil, errors.New("the grammar is mandatory in order to build an External instance")
	}

	return createExternal(app.name, app.grammar), nil
}
