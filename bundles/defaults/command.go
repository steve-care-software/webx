package defaults

import (
	creates_command "github.com/steve-care-software/syntax/applications/engines/creates/commands"
	"github.com/steve-care-software/syntax/domain/syntax/commands"
)

type command struct {
	builder                       commands.Builder
	executionBuilder              commands.ExecutionBuilder
	attachmentBuilder             commands.AttachmentBuilder
	variableAssignmentBuilder     commands.VariableAssignmentBuilder
	parameterDeclarationBuilder   commands.ParameterDeclarationBuilder
	applicationDeclarationBuilder commands.ApplicationDeclarationBuilder
	moduleDeclarationBuilder      commands.ModuleDeclarationBuilder
}

func createCommand(
	builder commands.Builder,
	executionBuilder commands.ExecutionBuilder,
	attachmentBuilder commands.AttachmentBuilder,
	variableAssignmentBuilder commands.VariableAssignmentBuilder,
	parameterDeclarationBuilder commands.ParameterDeclarationBuilder,
	applicationDeclarationBuilder commands.ApplicationDeclarationBuilder,
	moduleDeclarationBuilder commands.ModuleDeclarationBuilder,
) creates_command.Application {
	out := command{
		builder:                       builder,
		executionBuilder:              executionBuilder,
		attachmentBuilder:             attachmentBuilder,
		variableAssignmentBuilder:     variableAssignmentBuilder,
		parameterDeclarationBuilder:   parameterDeclarationBuilder,
		applicationDeclarationBuilder: applicationDeclarationBuilder,
		moduleDeclarationBuilder:      moduleDeclarationBuilder,
	}

	return &out
}

// Execute executes the application
func (app *command) Execute() (commands.Command, error) {
	return nil, nil
}

func (app *command) createCommand() (commands.Command, error) {
	return nil, nil
}

func (app *command) createCommandExecution() (commands.Execution, error) {
	return nil, nil
}

func (app *command) createCommandAttachment() (commands.Attachment, error) {
	return nil, nil
}

func (app *command) createCommandVariableAssignment() (commands.VariableAssignment, error) {
	return nil, nil
}

func (app *command) createCommandParameterDeclaration() (commands.ParameterDeclaration, error) {
	return nil, nil
}

func (app *command) createCommandApplicationDeclaration() (commands.ApplicationDeclaration, error) {
	return nil, nil
}

func (app *command) createCommandModuleDeclaration() (commands.ModuleDeclaration, error) {
	return nil, nil
}
