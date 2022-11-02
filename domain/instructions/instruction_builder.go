package instructions

import (
	"errors"

	"github.com/steve-care-software/webx/domain/instructions/applications"
	"github.com/steve-care-software/webx/domain/instructions/attachments"
	"github.com/steve-care-software/webx/domain/instructions/parameters"
)

type instructionBuilder struct {
	module      []byte
	application applications.Application
	parameter   parameters.Parameter
	assignment  Assignment
	attachment  attachments.Attachment
	execution   []byte
}

func createInstructionBuilder() InstructionBuilder {
	out := instructionBuilder{
		module:      nil,
		application: nil,
		parameter:   nil,
		assignment:  nil,
		attachment:  nil,
		execution:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *instructionBuilder) Create() InstructionBuilder {
	return createInstructionBuilder()
}

// WithModule adds a module to the builder
func (app *instructionBuilder) WithModule(module []byte) InstructionBuilder {
	app.module = module
	return app
}

// WithApplication adds an application to the builder
func (app *instructionBuilder) WithApplication(application applications.Application) InstructionBuilder {
	app.application = application
	return app
}

// WithParameter adds a parameter to the builder
func (app *instructionBuilder) WithParameter(parameter parameters.Parameter) InstructionBuilder {
	app.parameter = parameter
	return app
}

// WithAssignment adds an assignment to the builder
func (app *instructionBuilder) WithAssignment(assignment Assignment) InstructionBuilder {
	app.assignment = assignment
	return app
}

// WithAttachment adds an attachment to the builder
func (app *instructionBuilder) WithAttachment(attachment attachments.Attachment) InstructionBuilder {
	app.attachment = attachment
	return app
}

// WithExecution adds an execution to the builder
func (app *instructionBuilder) WithExecution(execution []byte) InstructionBuilder {
	app.execution = execution
	return app
}

// Now builds a new Instruction instance
func (app *instructionBuilder) Now() (Instruction, error) {
	if app.module != nil {
		return createInstructionWithModule(app.module), nil
	}

	if app.application != nil {
		return createInstructionWithApplication(app.application), nil
	}

	if app.parameter != nil {
		return createInstructionWithParameter(app.parameter), nil
	}

	if app.assignment != nil {
		return createInstructionWithAssignment(app.assignment), nil
	}

	if app.attachment != nil {
		return createInstructionWithAttachment(app.attachment), nil
	}

	if app.execution != nil {
		return createInstructionWithExecution(app.execution), nil
	}

	return nil, errors.New("the Instruction is invalid")
}
