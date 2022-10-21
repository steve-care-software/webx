package compilers

import (
	"errors"

	"github.com/steve-care-software/webx/domain/grammars"
	"github.com/steve-care-software/webx/domain/programs"
)

type elementBuilder struct {
	grammar    grammars.Grammar
	program    programs.Program
	parameters Parameters
}

func createElementBuilder() ElementBuilder {
	out := elementBuilder{
		grammar:    nil,
		program:    nil,
		parameters: nil,
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

// WithProgram adds a program to the builder
func (app *elementBuilder) WithProgram(program programs.Program) ElementBuilder {
	app.program = program
	return app
}

// WithParameters add parameters to the builder
func (app *elementBuilder) WithParameters(parameters Parameters) ElementBuilder {
	app.parameters = parameters
	return app
}

// Now builds a new Element instance
func (app *elementBuilder) Now() (Element, error) {
	if app.grammar == nil {
		return nil, errors.New("the grammar is mandatory in order to build an Element instance")
	}

	if app.program == nil {
		return nil, errors.New("the program is mandatory in order to build an Element instance")
	}

	if app.parameters == nil {
		return nil, errors.New("the parameters is mandatory in order to build an Element instance")
	}

	return createElement(app.grammar, app.program, app.parameters), nil
}
