package programs

import (
	"errors"

	"github.com/steve-care-software/syntax/domain/syntax/programs/instructions"
)

type builder struct {
	instructions instructions.Instructions
	inputs       []string
	outputs      []string
}

func createBuilder() Builder {
	out := builder{
		instructions: nil,
		inputs:       nil,
		outputs:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithInstructions add instructions to the builder
func (app *builder) WithInstructions(instructions instructions.Instructions) Builder {
	app.instructions = instructions
	return app
}

// WithInputs add inputs to the builder
func (app *builder) WithInputs(inputs []string) Builder {
	app.inputs = inputs
	return app
}

// WithOutputs add outputs to the builder
func (app *builder) WithOutputs(outputs []string) Builder {
	app.outputs = outputs
	return app
}

// Now builds a new Program instance
func (app *builder) Now() (Program, error) {
	if app.instructions == nil {
		return nil, errors.New("the instructions is mandatory in order to build a Program instance")
	}

	if app.inputs != nil && len(app.inputs) <= 0 {
		app.inputs = nil
	}

	if app.outputs != nil && len(app.outputs) <= 0 {
		app.outputs = nil
	}

	if app.inputs != nil && app.outputs != nil {
		return createProgramWithInputsAndOutputs(app.instructions, app.inputs, app.outputs), nil
	}

	if app.inputs != nil {
		return createProgramWithInputs(app.instructions, app.inputs), nil
	}

	if app.outputs != nil {
		return createProgramWithOutputs(app.instructions, app.outputs), nil
	}

	return createProgram(app.instructions), nil
}
