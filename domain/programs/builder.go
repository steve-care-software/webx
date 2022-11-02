package programs

import (
	"errors"
)

type builder struct {
	instructions []Instruction
	outputs      [][]byte
}

func createBuilder() Builder {
	out := builder{
		instructions: nil,
		outputs:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithInstructions add instructions to the builder
func (app *builder) WithInstructions(instructions []Instruction) Builder {
	app.instructions = instructions
	return app
}

// WithOutputs add outputs to the builder
func (app *builder) WithOutputs(outputs [][]byte) Builder {
	app.outputs = outputs
	return app
}

// Now builds a new Program instance
func (app *builder) Now() (Program, error) {
	if app.instructions != nil && len(app.instructions) <= 0 {
		app.instructions = nil
	}

	if app.instructions == nil {
		return nil, errors.New("the instructions are mandatory in order to build a Program instance")
	}

	if app.outputs != nil && len(app.outputs) <= 0 {
		app.outputs = nil
	}

	if app.outputs != nil {
		return createProgramWithOutputs(app.instructions, app.outputs), nil
	}

	return createProgram(app.instructions), nil
}
