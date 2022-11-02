package compilers

import (
	"github.com/steve-care-software/webx/applications/creates"
	"github.com/steve-care-software/webx/applications/grammars"
	"github.com/steve-care-software/webx/applications/interpreters"
	"github.com/steve-care-software/webx/domain/compilers"
	"github.com/steve-care-software/webx/domain/instructions"
	"github.com/steve-care-software/webx/domain/instructions/applications"
	"github.com/steve-care-software/webx/domain/instructions/attachments"
	"github.com/steve-care-software/webx/domain/instructions/parameters"
)

// ExecuteInstructionsModuleName represents the execute instructions module name
const ExecuteInstructionsModuleName = "executeInstructions"

// NewApplication creates a new application
func NewApplication(
	createApp creates.Application,
) Application {
	grammarApp := grammars.NewApplication()
	interpreterApp := interpreters.NewApplication()
	builder := instructions.NewBuilder()
	instructionBuilder := instructions.NewInstructionBuilder()
	assignmentBuilder := instructions.NewAssignmentBuilder()
	valueBuilder := instructions.NewValueBuilder()
	applicationBuilder := applications.NewBuilder()
	attachmentBuilder := attachments.NewBuilder()
	attachmentVariableBuilder := attachments.NewVariableBuilder()
	parameterBuilder := parameters.NewBuilder()
	return createApplication(
		grammarApp,
		interpreterApp,
		createApp,
		builder,
		instructionBuilder,
		assignmentBuilder,
		valueBuilder,
		applicationBuilder,
		attachmentBuilder,
		attachmentVariableBuilder,
		parameterBuilder,
	)
}

// Application represents the compiler application
type Application interface {
	Execute(compiler compilers.Compiler, script []byte) (instructions.Instructions, error)
}
