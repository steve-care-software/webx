package instructions

import (
	"github.com/steve-care-software/webx/domain/instructions/applications"
	"github.com/steve-care-software/webx/domain/instructions/attachments"
	"github.com/steve-care-software/webx/domain/instructions/parameters"
)

type instruction struct {
	module      []byte
	application applications.Application
	parameter   parameters.Parameter
	assignment  Assignment
	attachment  attachments.Attachment
	execution   []byte
}

func createInstructionWithModule(
	module []byte,
) Instruction {
	return createInstructionInternally(module, nil, nil, nil, nil, nil)
}

func createInstructionWithApplication(
	application applications.Application,
) Instruction {
	return createInstructionInternally(nil, application, nil, nil, nil, nil)
}

func createInstructionWithParameter(
	parameter parameters.Parameter,
) Instruction {
	return createInstructionInternally(nil, nil, parameter, nil, nil, nil)
}

func createInstructionWithAssignment(
	assignment Assignment,
) Instruction {
	return createInstructionInternally(nil, nil, nil, assignment, nil, nil)
}

func createInstructionWithAttachment(
	attachment attachments.Attachment,
) Instruction {
	return createInstructionInternally(nil, nil, nil, nil, attachment, nil)
}

func createInstructionWithExecution(
	execution []byte,
) Instruction {
	return createInstructionInternally(nil, nil, nil, nil, nil, execution)
}

func createInstructionInternally(
	module []byte,
	application applications.Application,
	parameter parameters.Parameter,
	assignment Assignment,
	attachment attachments.Attachment,
	execution []byte,
) Instruction {
	out := instruction{
		module:      module,
		application: application,
		parameter:   parameter,
		assignment:  assignment,
		attachment:  attachment,
		execution:   execution,
	}

	return &out
}

// IsModule returns true if there is a module, false otherwise
func (obj *instruction) IsModule() bool {
	return obj.module != nil
}

// Module returns the module, if any
func (obj *instruction) Module() []byte {
	return obj.module
}

// IsApplication returns true if there is an application, false otherwise
func (obj *instruction) IsApplication() bool {
	return obj.application != nil
}

// Application returns the application, if any
func (obj *instruction) Application() applications.Application {
	return obj.application
}

// IsParameter returns true if there is a parameter, false otherwise
func (obj *instruction) IsParameter() bool {
	return obj.parameter != nil
}

// Parameter returns the parameter, if any
func (obj *instruction) Parameter() parameters.Parameter {
	return obj.parameter
}

// IsAssignment returns true if there is an assignment, false otherwise
func (obj *instruction) IsAssignment() bool {
	return obj.assignment != nil
}

// Assignment returns the assignment, if any
func (obj *instruction) Assignment() Assignment {
	return obj.assignment
}

// IsAttachment returns true if there is an attachment, false otherwise
func (obj *instruction) IsAttachment() bool {
	return obj.attachment != nil
}

// Attachment returns the attachment, if any
func (obj *instruction) Attachment() attachments.Attachment {
	return obj.attachment
}

// IsExecution returns true if there is an execution, false otherwise
func (obj *instruction) IsExecution() bool {
	return obj.execution != nil
}

// Execution returns the execution, if any
func (obj *instruction) Execution() []byte {
	return obj.execution
}
