package defaults

import (
	creates_command "github.com/steve-care-software/syntax/applications/engines/creates/commands"
	"github.com/steve-care-software/syntax/domain/syntax/bytes/criterias"
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
	criteriaBuilder               criterias.Builder
}

func createCommand(
	builder commands.Builder,
	executionBuilder commands.ExecutionBuilder,
	attachmentBuilder commands.AttachmentBuilder,
	variableAssignmentBuilder commands.VariableAssignmentBuilder,
	parameterDeclarationBuilder commands.ParameterDeclarationBuilder,
	applicationDeclarationBuilder commands.ApplicationDeclarationBuilder,
	moduleDeclarationBuilder commands.ModuleDeclarationBuilder,
	criteriaBuilder criterias.Builder,
) creates_command.Application {
	out := command{
		builder:                       builder,
		executionBuilder:              executionBuilder,
		attachmentBuilder:             attachmentBuilder,
		variableAssignmentBuilder:     variableAssignmentBuilder,
		parameterDeclarationBuilder:   parameterDeclarationBuilder,
		applicationDeclarationBuilder: applicationDeclarationBuilder,
		moduleDeclarationBuilder:      moduleDeclarationBuilder,
		criteriaBuilder:               criteriaBuilder,
	}

	return &out
}

// Execute executes the application
func (app *command) Execute() (commands.Command, error) {
	execution, err := app.execution()
	if err != nil {
		return nil, err
	}

	attachment, err := app.attachment()
	if err != nil {
		return nil, err
	}

	variableAssignment, err := app.variableAssignment()
	if err != nil {
		return nil, err
	}

	parameterDeclaration, err := app.parameterDeclaration()
	if err != nil {
		return nil, err
	}

	applicationDeclaration, err := app.applicationDeclaration()
	if err != nil {
		return nil, err
	}

	moduleDeclaration, err := app.moduleDeclaration()
	if err != nil {
		return nil, err
	}

	return app.builder.Create().
		WithExecution(execution).
		WithAttachment(attachment).
		WithVariableAssignment(variableAssignment).
		WithParameterDeclaration(parameterDeclaration).
		WithApplicationDeclaration(applicationDeclaration).
		WithModuleDeclaration(moduleDeclaration).
		Now()
}

func (app *command) execution() (commands.Execution, error) {
	assignee, err := app.executionAssignee()
	if err != nil {
		return nil, err
	}

	applicationName, err := app.executionApplicationName()
	if err != nil {
		return nil, err
	}

	return app.executionBuilder.Create().WithAssignee(assignee).WithApplication(applicationName).Now()
}

func (app *command) executionAssignee() (criterias.Criteria, error) {
	variableName, err := app.variableName(0)
	if err != nil {
		return nil, err
	}

	executeAssignment, err := app.criteriaBuilder.Create().WithName("executeAssignment").WithIndex(0).WithChild(variableName).Now()
	if err != nil {
		return nil, err
	}

	return app.criteriaBuilder.Create().WithName("root").WithIndex(0).WithChild(executeAssignment).Now()
}

func (app *command) executionApplicationName() (criterias.Criteria, error) {
	variableName, err := app.variableName(0)
	if err != nil {
		return nil, err
	}

	execute, err := app.criteriaBuilder.Create().WithName("execute").WithIndex(0).WithChild(variableName).Now()
	if err != nil {
		return nil, err
	}

	executeAssignment, err := app.criteriaBuilder.Create().WithName("executeAssignment").WithIndex(0).WithChild(execute).Now()
	if err != nil {
		return nil, err
	}

	return app.criteriaBuilder.Create().WithName("root").WithIndex(0).WithChild(executeAssignment).Now()
}

func (app *command) attachment() (commands.Attachment, error) {
	global, err := app.attachmentVariableName(0)
	if err != nil {
		return nil, err
	}

	local, err := app.attachmentVariableName(1)
	if err != nil {
		return nil, err
	}

	application, err := app.attachmentVariableName(2)
	if err != nil {
		return nil, err
	}

	return app.attachmentBuilder.Create().WithGlobal(global).WithLocal(local).WithApplication(application).Now()
}

func (app *command) attachmentVariableName(variableIndex uint) (criterias.Criteria, error) {
	variableName, err := app.variableName(variableIndex)
	if err != nil {
		return nil, err
	}

	attachment, err := app.criteriaBuilder.Create().WithName("attachment").WithIndex(0).WithChild(variableName).Now()
	if err != nil {
		return nil, err
	}

	return app.criteriaBuilder.Create().WithName("root").WithIndex(0).WithChild(attachment).Now()
}

func (app *command) variableAssignment() (commands.VariableAssignment, error) {
	value, err := app.variableAssignmentValue()
	if err != nil {
		return nil, err
	}

	assignee, err := app.variableAssignmentAssignee()
	if err != nil {
		return nil, err
	}

	return app.variableAssignmentBuilder.Create().WithValue(value).WithAssignee(assignee).Now()
}

func (app *command) variableAssignmentAssignee() (criterias.Criteria, error) {
	variableName, err := app.variableName(0)
	if err != nil {
		return nil, err
	}

	assignment, err := app.criteriaBuilder.Create().WithName("variableAssignment").WithIndex(0).WithChild(variableName).Now()
	if err != nil {
		return nil, err
	}

	return app.criteriaBuilder.Create().WithName("root").WithIndex(0).WithChild(assignment).Now()
}

func (app *command) variableAssignmentValue() (criterias.Criteria, error) {
	assignee, err := app.criteriaBuilder.Create().WithName("assignmentValue").WithIndex(0).Now()
	if err != nil {
		return nil, err
	}

	assignment, err := app.criteriaBuilder.Create().WithName("variableAssignment").WithIndex(0).WithChild(assignee).Now()
	if err != nil {
		return nil, err
	}

	return app.criteriaBuilder.Create().WithName("root").WithIndex(0).WithChild(assignment).Now()
}

func (app *command) parameterDeclaration() (commands.ParameterDeclaration, error) {
	input, err := app.parameterDeclarationInput()
	if err != nil {
		return nil, err
	}

	output, err := app.parameterDeclarationOutput()
	if err != nil {
		return nil, err
	}

	return app.parameterDeclarationBuilder.Create().WithInput(input).WithOutput(output).Now()
}

func (app *command) parameterDeclarationInput() (criterias.Criteria, error) {
	variableName, err := app.variableName(0)
	if err != nil {
		return nil, err
	}

	inputParameter, err := app.criteriaBuilder.Create().WithName("inputParameter").WithIndex(0).WithChild(variableName).Now()
	if err != nil {
		return nil, err
	}

	return app.criteriaBuilder.Create().WithName("root").WithIndex(0).WithChild(inputParameter).Now()
}

func (app *command) parameterDeclarationOutput() (criterias.Criteria, error) {
	variableName, err := app.variableName(0)
	if err != nil {
		return nil, err
	}

	outputParameter, err := app.criteriaBuilder.Create().WithName("outputParameter").WithIndex(0).WithChild(variableName).Now()
	if err != nil {
		return nil, err
	}

	return app.criteriaBuilder.Create().WithName("root").WithIndex(0).WithChild(outputParameter).Now()
}

func (app *command) applicationDeclaration() (commands.ApplicationDeclaration, error) {
	module, err := app.applicationDeclarationModule()
	if err != nil {
		return nil, err
	}

	name, err := app.applicationDeclarationName()
	if err != nil {
		return nil, err
	}

	return app.applicationDeclarationBuilder.Create().WithModule(module).WithName(name).Now()
}

func (app *command) applicationDeclarationModule() (criterias.Criteria, error) {
	name, err := app.criteriaBuilder.Create().WithName("name").WithIndex(0).Now()
	if err != nil {
		return nil, err
	}

	applicationDeclaration, err := app.criteriaBuilder.Create().WithName("applicationDeclaration").WithIndex(0).WithChild(name).Now()
	if err != nil {
		return nil, err
	}

	return app.criteriaBuilder.Create().WithName("root").WithIndex(0).WithChild(applicationDeclaration).Now()
}

func (app *command) applicationDeclarationName() (criterias.Criteria, error) {
	variableName, err := app.variableName(0)
	if err != nil {
		return nil, err
	}

	applicationDeclaration, err := app.criteriaBuilder.Create().WithName("applicationDeclaration").WithIndex(0).WithChild(variableName).Now()
	if err != nil {
		return nil, err
	}

	return app.criteriaBuilder.Create().WithName("root").WithIndex(0).WithChild(applicationDeclaration).Now()
}

func (app *command) moduleDeclaration() (commands.ModuleDeclaration, error) {
	name, err := app.criteriaBuilder.Create().WithName("name").WithIndex(0).Now()
	if err != nil {
		return nil, err
	}

	moduleDeclaration, err := app.criteriaBuilder.Create().WithName("moduleDeclaration").WithIndex(0).WithChild(name).Now()
	if err != nil {
		return nil, err
	}

	root, err := app.criteriaBuilder.Create().WithName("root").WithIndex(0).WithChild(moduleDeclaration).Now()
	if err != nil {
		return nil, err
	}

	return app.moduleDeclarationBuilder.Create().WithName(root).Now()
}

func (app *command) variableName(index uint) (criterias.Criteria, error) {
	name, err := app.criteriaBuilder.Create().WithName("name").WithIndex(0).Now()
	if err != nil {
		return nil, err
	}

	return app.criteriaBuilder.Create().WithName("variableName").WithIndex(index).WithChild(name).Now()
}
