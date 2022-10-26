package defaults

import (
	creates_command "github.com/steve-care-software/webx/applications/creates/commands"
	"github.com/steve-care-software/webx/domain/commands"
	"github.com/steve-care-software/webx/domain/criterias"
)

type command struct {
	builder                       commands.Builder
	attachmentBuilder             commands.AttachmentBuilder
	variableAssignmentBuilder     commands.VariableAssignmentBuilder
	parameterDeclarationBuilder   commands.ParameterDeclarationBuilder
	applicationDeclarationBuilder commands.ApplicationDeclarationBuilder
	valueBuilder                  commands.ValueBuilder
	criteriaBuilder               criterias.Builder
}

func createCommand(
	builder commands.Builder,
	attachmentBuilder commands.AttachmentBuilder,
	variableAssignmentBuilder commands.VariableAssignmentBuilder,
	parameterDeclarationBuilder commands.ParameterDeclarationBuilder,
	applicationDeclarationBuilder commands.ApplicationDeclarationBuilder,
	valueBuilder commands.ValueBuilder,
	criteriaBuilder criterias.Builder,
) creates_command.Application {
	out := command{
		builder:                       builder,
		attachmentBuilder:             attachmentBuilder,
		variableAssignmentBuilder:     variableAssignmentBuilder,
		parameterDeclarationBuilder:   parameterDeclarationBuilder,
		applicationDeclarationBuilder: applicationDeclarationBuilder,
		valueBuilder:                  valueBuilder,
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

func (app *command) execution() (criterias.Criteria, error) {
	variableName := app.VariableName(0)
	execute, err := app.criteriaBuilder.Create().WithName("execute").WithIndex(0).WithChild(variableName).Now()
	if err != nil {
		return nil, err
	}

	executeAssignment, err := app.criteriaBuilder.Create().WithName("executeAssignment").WithIndex(0).WithChild(execute).Now()
	if err != nil {
		return nil, err
	}

	return app.criteriaBuilder.Create().WithName("instruction").WithIndex(0).WithChild(executeAssignment).Now()
}

func (app *command) executionAssignee() (criterias.Criteria, error) {
	variableName := app.VariableName(0)
	executeAssignment, err := app.criteriaBuilder.Create().WithName("executeAssignment").WithIndex(0).WithChild(variableName).Now()
	if err != nil {
		return nil, err
	}

	return app.criteriaBuilder.Create().WithName("instruction").WithIndex(0).WithChild(executeAssignment).Now()
}

func (app *command) attachment() (commands.Attachment, error) {
	current, err := app.attachmentVariableName(0)
	if err != nil {
		return nil, err
	}

	target, err := app.attachmentVariableName(1)
	if err != nil {
		return nil, err
	}

	application, err := app.attachmentVariableName(2)
	if err != nil {
		return nil, err
	}

	return app.attachmentBuilder.Create().WithCurrent(current).WithTarget(target).WithApplication(application).Now()
}

func (app *command) attachmentVariableName(variableIndex uint) (criterias.Criteria, error) {
	variableName := app.VariableName(variableIndex)
	attachment, err := app.criteriaBuilder.Create().WithName("attachment").WithIndex(0).WithChild(variableName).Now()
	if err != nil {
		return nil, err
	}

	return app.criteriaBuilder.Create().WithName("instruction").WithIndex(0).WithChild(attachment).Now()
}

func (app *command) variableAssignment() (commands.VariableAssignment, error) {
	value, err := app.value()
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
	variableName := app.VariableName(0)
	assignment, err := app.criteriaBuilder.Create().WithName("assignment").WithIndex(0).WithChild(variableName).Now()
	if err != nil {
		return nil, err
	}

	return app.criteriaBuilder.Create().WithName("instruction").WithIndex(0).WithChild(assignment).Now()
}

func (app *command) value() (commands.Value, error) {
	variable, err := app.valueVariable()
	if err != nil {
		return nil, err
	}

	constant, err := app.valueConstant()
	if err != nil {
		return nil, err
	}

	instructions, err := app.valueInstructions()
	if err != nil {
		return nil, err
	}

	execution, err := app.valueExecution()
	if err != nil {
		return nil, err
	}

	return app.valueBuilder.Create().WithVariable(variable).
		WithConstant(constant).
		WithInstructions(instructions).
		WithExecution(execution).
		Now()
}

func (app *command) valueVariable() (criterias.Criteria, error) {
	value, err := app.criteriaBuilder.Create().WithName("valueVariable").WithIndex(0).Now()
	if err != nil {
		return nil, err
	}

	assignment, err := app.criteriaBuilder.Create().WithName("assignment").WithIndex(0).WithChild(value).Now()
	if err != nil {
		return nil, err
	}

	return app.criteriaBuilder.Create().WithName("instruction").WithIndex(0).WithChild(assignment).Now()
}

func (app *command) valueConstant() (criterias.Criteria, error) {
	value, err := app.criteriaBuilder.Create().WithName("assignmentValue").WithIndex(0).Now()
	if err != nil {
		return nil, err
	}

	valueConstant, err := app.criteriaBuilder.Create().WithName("valueConstant").WithIndex(0).WithChild(value).Now()
	if err != nil {
		return nil, err
	}

	assignment, err := app.criteriaBuilder.Create().WithName("assignment").WithIndex(0).WithChild(valueConstant).Now()
	if err != nil {
		return nil, err
	}

	return app.criteriaBuilder.Create().WithName("instruction").WithIndex(0).WithChild(assignment).Now()
}

func (app *command) valueInstructions() (criterias.Criteria, error) {
	root, err := app.criteriaBuilder.Create().WithName("instruction").WithIndex(0).Now()
	if err != nil {
		return nil, err
	}

	value, err := app.criteriaBuilder.Create().WithName("codeAssignment").WithChild(root).WithIndex(0).Now()
	if err != nil {
		return nil, err
	}

	assignment, err := app.criteriaBuilder.Create().WithName("assignment").WithIndex(0).WithChild(value).Now()
	if err != nil {
		return nil, err
	}

	return app.criteriaBuilder.Create().WithName("instruction").WithIndex(0).WithChild(assignment).Now()
}

func (app *command) valueExecution() (criterias.Criteria, error) {
	variableName := app.VariableName(0)
	execute, err := app.criteriaBuilder.Create().WithName("execute").WithIndex(0).WithChild(variableName).Now()
	if err != nil {
		return nil, err
	}

	value, err := app.criteriaBuilder.Create().WithName("valueExecution").WithIndex(0).WithChild(execute).Now()
	if err != nil {
		return nil, err
	}

	assignment, err := app.criteriaBuilder.Create().WithName("assignment").WithIndex(0).WithChild(value).Now()
	if err != nil {
		return nil, err
	}

	return app.criteriaBuilder.Create().WithName("instruction").WithIndex(0).WithChild(assignment).Now()
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
	variableName := app.VariableName(0)
	inputParameter, err := app.criteriaBuilder.Create().WithName("inputParameter").WithIndex(0).WithChild(variableName).Now()
	if err != nil {
		return nil, err
	}

	return app.criteriaBuilder.Create().WithName("instruction").WithIndex(0).WithChild(inputParameter).Now()
}

func (app *command) parameterDeclarationOutput() (criterias.Criteria, error) {
	variableName := app.VariableName(0)
	outputParameter, err := app.criteriaBuilder.Create().WithName("outputParameter").WithIndex(0).WithChild(variableName).Now()
	if err != nil {
		return nil, err
	}

	return app.criteriaBuilder.Create().WithName("instruction").WithIndex(0).WithChild(outputParameter).Now()
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

	return app.criteriaBuilder.Create().WithName("instruction").WithIndex(0).WithChild(applicationDeclaration).Now()
}

func (app *command) applicationDeclarationName() (criterias.Criteria, error) {
	variableName := app.VariableName(0)
	applicationDeclaration, err := app.criteriaBuilder.Create().WithName("applicationDeclaration").WithIndex(0).WithChild(variableName).Now()
	if err != nil {
		return nil, err
	}

	return app.criteriaBuilder.Create().WithName("instruction").WithIndex(0).WithChild(applicationDeclaration).Now()
}

func (app *command) moduleDeclaration() (criterias.Criteria, error) {
	name, err := app.criteriaBuilder.Create().WithName("name").WithIndex(0).Now()
	if err != nil {
		return nil, err
	}

	moduleDeclaration, err := app.criteriaBuilder.Create().WithName("moduleDeclaration").WithIndex(0).WithChild(name).Now()
	if err != nil {
		return nil, err
	}

	return app.criteriaBuilder.Create().WithName("instruction").WithIndex(0).WithChild(moduleDeclaration).Now()
}

// VariableName returns the variable name criteria
func (app *command) VariableName(index uint) criterias.Criteria {
	name, err := app.criteriaBuilder.Create().WithName("name").WithIndex(0).Now()
	if err != nil {
		panic(err)
	}

	ins, err := app.criteriaBuilder.Create().WithName("variableName").WithIndex(index).WithChild(name).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
