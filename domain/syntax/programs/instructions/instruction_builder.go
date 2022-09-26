package instructions

import (
	"errors"

	"github.com/steve-care-software/syntax/domain/syntax/programs/instructions/applications"
)

type instructionBuilder struct {
	assignment applications.Assignment
	execution  applications.Application
	delete     applications.Application
	setPath    string
}

func createInstructionBuilder() InstructionBuilder {
	out := instructionBuilder{
		assignment: nil,
		execution:  nil,
		delete:     nil,
		setPath:    "",
	}

	return &out
}

// Create initializes the builder
func (app *instructionBuilder) Create() InstructionBuilder {
	return createInstructionBuilder()
}

// WithAssignment adds an assignment to the builder
func (app *instructionBuilder) WithAssignment(assignment applications.Assignment) InstructionBuilder {
	app.assignment = assignment
	return app
}

// WithExecution adds an execution to the builder
func (app *instructionBuilder) WithExecution(execution applications.Application) InstructionBuilder {
	app.execution = execution
	return app
}

// WithDelete adds a delete to the builder
func (app *instructionBuilder) WithDelete(delete applications.Application) InstructionBuilder {
	app.delete = delete
	return app
}

// WithSetPath adds a setPath to the builder
func (app *instructionBuilder) WithSetPath(setPath string) InstructionBuilder {
	app.setPath = setPath
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

	if app.delete != nil {
		return createInstructionWithDelete(app.delete), nil
	}

	if app.setPath != "" {
		return createInstructionWithSetPath(app.setPath), nil
	}

	return nil, errors.New("the Instruction is invalid")
}
