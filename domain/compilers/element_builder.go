package compilers

import (
	"errors"

	"github.com/steve-care-software/webx/domain/grammars"
)

type elementBuilder struct {
	grammar    grammars.Grammar
	execution  Execution
	parameters Parameters
}

func createElementBuilder() ElementBuilder {
	out := elementBuilder{
		grammar:    nil,
		execution:  nil,
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

// WithExecution adds an execution to the builder
func (app *elementBuilder) WithExecution(execution Execution) ElementBuilder {
	app.execution = execution
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

	if app.execution == nil {
		return nil, errors.New("the execution is mandatory in order to build an Element instance")
	}

	if app.parameters == nil {
		return nil, errors.New("the parameters is mandatory in order to build an Element instance")
	}

	return createElement(app.grammar, app.execution, app.parameters), nil
}
