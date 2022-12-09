package compilers

import (
	"errors"

	"github.com/steve-care-software/webx/programs/domain/instructions"
)

type executionBuilder struct {
	parameter    string
	instructions instructions.Instructions
}

func createExecutionBuilder() ExecutionBuilder {
	out := executionBuilder{
		parameter:    "",
		instructions: nil,
	}

	return &out
}

// Create initializes the builder
func (app *executionBuilder) Create() ExecutionBuilder {
	return createExecutionBuilder()
}

// WithParameter adds a parameter to the builder
func (app *executionBuilder) WithParameter(parameter string) ExecutionBuilder {
	app.parameter = parameter
	return app
}

// WithInstructions add instructions to the builder
func (app *executionBuilder) WithInstructions(instructions instructions.Instructions) ExecutionBuilder {
	app.instructions = instructions
	return app
}

// Now builds a new Execution instance
func (app *executionBuilder) Now() (Execution, error) {
	if app.parameter == "" {
		return nil, errors.New("the parameter is mandatory in order to build an Execution instance")
	}

	if app.instructions == nil {
		return nil, errors.New("the instructions is mandatory in order to build an Execution instance")
	}

	return createExecution(app.parameter, app.instructions), nil
}
