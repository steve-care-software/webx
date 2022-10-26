package compilers

import (
	"github.com/steve-care-software/webx/applications/creates"
	"github.com/steve-care-software/webx/applications/grammars"
	"github.com/steve-care-software/webx/applications/interpreters"
	"github.com/steve-care-software/webx/domain/compilers"
	"github.com/steve-care-software/webx/domain/instructions"
	"github.com/steve-care-software/webx/domain/instructions/applications"
	"github.com/steve-care-software/webx/domain/instructions/attachments"
)

// ExecuteInstructionsModuleName represents the execute instructions module name
const ExecuteInstructionsModuleName = "executeInstructions"

// NewApplication creates a new application
func NewApplication(
	createApp creates.Application,
) Application {
	grammarApp := grammars.NewApplication()
	//criteriaApp := criterias.NewApplication()
	interpreterApp := interpreters.NewApplication()
	instructionsBuilder := instructions.NewBuilder()
	instructionBuilder := instructions.NewInstructionBuilder()
	applicationBuilder := applications.NewBuilder()
	attachmentBuilder := attachments.NewBuilder()
	attachmentVariableBuilder := attachments.NewVariableBuilder()
	assignmmentBuilder := instructions.NewAssignmentBuilder()
	valueBuilder := instructions.NewValueBuilder()
	outputBuilder := instructions.NewOutputBuilder()
	return createApplication(
		grammarApp,
		//criteriaApp,
		interpreterApp,
		createApp,
		instructionsBuilder,
		instructionBuilder,
		applicationBuilder,
		attachmentBuilder,
		attachmentVariableBuilder,
		assignmmentBuilder,
		valueBuilder,
		outputBuilder,
	)
}

// Application represents the compiler application
type Application interface {
	Execute(compiler compilers.Compiler, script []byte) (instructions.Output, error)
}
