package commands

import (
	"errors"

	"github.com/steve-care-software/webx/domain/criterias"
)

type valueBuilder struct {
	variable     criterias.Criteria
	constant     criterias.Criteria
	instructions criterias.Criteria
	execution    criterias.Criteria
}

func createValueBuilder() ValueBuilder {
	out := valueBuilder{
		variable:     nil,
		constant:     nil,
		instructions: nil,
		execution:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *valueBuilder) Create() ValueBuilder {
	return createValueBuilder()
}

// WithVariable adds a variable to the builder
func (app *valueBuilder) WithVariable(variable criterias.Criteria) ValueBuilder {
	app.variable = variable
	return app
}

// WithConstant adds a constant to the builder
func (app *valueBuilder) WithConstant(constant criterias.Criteria) ValueBuilder {
	app.constant = constant
	return app
}

// WithInstructions adds an instructions to the builder
func (app *valueBuilder) WithInstructions(instructions criterias.Criteria) ValueBuilder {
	app.instructions = instructions
	return app
}

// WithExecution adds an execution to the builder
func (app *valueBuilder) WithExecution(execution criterias.Criteria) ValueBuilder {
	app.execution = execution
	return app
}

// Now builds a new Value instance
func (app *valueBuilder) Now() (Value, error) {
	if app.variable == nil {
		return nil, errors.New("the variable is mandatory in order to build a Value instance")
	}

	if app.constant == nil {
		return nil, errors.New("the constant is mandatory in order to build a Value instance")
	}

	if app.instructions == nil {
		return nil, errors.New("the instructions is mandatory in order to build a Value instance")
	}

	if app.execution == nil {
		return nil, errors.New("the execution is mandatory in order to build a Value instance")
	}

	return createValue(app.variable, app.constant, app.instructions, app.execution), nil
}
