package programs

import "errors"

type instructionBuilder struct {
	assignment Assignment
	execution  Application
}

func createInstructionBuilder() InstructionBuilder {
	out := instructionBuilder{
		assignment: nil,
		execution:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *instructionBuilder) Create() InstructionBuilder {
	return createInstructionBuilder()
}

// WithAssignment adds an assignment to the builder
func (app *instructionBuilder) WithAssignment(assignment Assignment) InstructionBuilder {
	app.assignment = assignment
	return app
}

// WithExecution adds an execution to the builder
func (app *instructionBuilder) WithExecution(execution Application) InstructionBuilder {
	app.execution = execution
	return app
}

// Now builds a new Instruction instance
func (app *instructionBuilder) Now() (Instruction, error) {
	if app.assignment != nil {
		return createInstructionWithAssignment(app.assignment), nil
	}

	if app.execution != nil {
		return createInstructionWithExecution(app.execution), nil
	}

	return nil, errors.New("the Instruction is invalid")
}
