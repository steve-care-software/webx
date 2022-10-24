package instructions

import (
	"github.com/steve-care-software/webx/applications/criterias"
	grammar_application "github.com/steve-care-software/webx/applications/grammars"
	"github.com/steve-care-software/webx/domain/commands"
	"github.com/steve-care-software/webx/domain/grammars"
	"github.com/steve-care-software/webx/domain/instructions"
	"github.com/steve-care-software/webx/domain/instructions/applications"
	"github.com/steve-care-software/webx/domain/instructions/attachments"
	"github.com/steve-care-software/webx/domain/instructions/parameters"
)

// NewApplication creates a new application
func NewApplication() Application {
	grammarApp := grammar_application.NewApplication()
	criteriaApp := criterias.NewApplication()
	builder := instructions.NewBuilder()
	instructionBuilder := instructions.NewInstructionBuilder()
	assignmentBuilder := instructions.NewAssignmentBuilder()
	valueBuilder := instructions.NewValueBuilder()
	applicationBuilder := applications.NewBuilder()
	attachmentBuilder := attachments.NewBuilder()
	attachmentVariableBuilder := attachments.NewVariableBuilder()
	parameterBuilder := parameters.NewBuilder()
	outputBuilder := instructions.NewOutputBuilder()
	return createApplication(
		grammarApp,
		criteriaApp,
		builder,
		instructionBuilder,
		assignmentBuilder,
		valueBuilder,
		applicationBuilder,
		attachmentBuilder,
		attachmentVariableBuilder,
		parameterBuilder,
		outputBuilder,
	)
}

// Application represents an instruction application
type Application interface {
	Execute(grammar grammars.Grammar, command commands.Command, script []byte) (instructions.Output, error)
}
