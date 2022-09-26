package commands

import (
	"errors"

	"github.com/steve-care-software/syntax/domain/bytes/criterias"
)

type parameterDeclarationBuilder struct {
	input  criterias.Criteria
	output criterias.Criteria
	name   criterias.Criteria
}

func createParameterDeclarationBuilder() ParameterDeclarationBuilder {
	out := parameterDeclarationBuilder{
		input:  nil,
		output: nil,
		name:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *parameterDeclarationBuilder) Create() ParameterDeclarationBuilder {
	return createParameterDeclarationBuilder()
}

// WithInput adds an input to the builder
func (app *parameterDeclarationBuilder) WithInput(input criterias.Criteria) ParameterDeclarationBuilder {
	app.input = input
	return app
}

// WithOutput adds an output to the builder
func (app *parameterDeclarationBuilder) WithOutput(output criterias.Criteria) ParameterDeclarationBuilder {
	app.output = output
	return app
}

// WithName adds a name to the builder
func (app *parameterDeclarationBuilder) WithName(name criterias.Criteria) ParameterDeclarationBuilder {
	app.name = name
	return app
}

// Now builds a new ParameterDeclaration instance
func (app *parameterDeclarationBuilder) Now() (ParameterDeclaration, error) {
	if app.input == nil {
		return nil, errors.New("the input is mandatory in order to build a ParameterDeclaration instance")
	}

	if app.output == nil {
		return nil, errors.New("the output is mandatory in order to build a ParameterDeclaration instance")
	}

	if app.name == nil {
		return nil, errors.New("the name is mandatory in order to build a ParameterDeclaration instance")
	}

	return createParameterDeclaration(app.input, app.output, app.name), nil
}
