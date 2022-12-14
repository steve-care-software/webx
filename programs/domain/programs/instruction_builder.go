package programs

import "errors"

type instructionBuilder struct {
	value     Value
	execution Application
}

func createInstructionBuilder() InstructionBuilder {
	out := instructionBuilder{
		value:     nil,
		execution: nil,
	}

	return &out
}

// Create initializes the builder
func (app *instructionBuilder) Create() InstructionBuilder {
	return createInstructionBuilder()
}

// WithValue adds a value to the builder
func (app *instructionBuilder) WithValue(value Value) InstructionBuilder {
	app.value = value
	return app
}

// WithExecution adds an execution to the builder
func (app *instructionBuilder) WithExecution(execution Application) InstructionBuilder {
	app.execution = execution
	return app
}

// Now builds a new Instruction instance
func (app *instructionBuilder) Now() (Instruction, error) {
	if app.value != nil {
		return createInstructionWithValue(app.value), nil
	}

	if app.execution != nil {
		return createInstructionWithExecution(app.execution), nil
	}

	return nil, errors.New("the Instruction is invalid")
}
