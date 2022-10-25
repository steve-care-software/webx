package instructions

import (
	"errors"
	"fmt"

	application_criteria "github.com/steve-care-software/webx/applications/criterias"
	grammar_application "github.com/steve-care-software/webx/applications/grammars"
	"github.com/steve-care-software/webx/domain/commands"
	"github.com/steve-care-software/webx/domain/criterias"
	"github.com/steve-care-software/webx/domain/grammars"
	"github.com/steve-care-software/webx/domain/instructions"
	"github.com/steve-care-software/webx/domain/instructions/applications"
	"github.com/steve-care-software/webx/domain/instructions/attachments"
	"github.com/steve-care-software/webx/domain/instructions/parameters"
	"github.com/steve-care-software/webx/domain/trees"
)

type application struct {
	grammarApp                grammar_application.Application
	criteriaApp               application_criteria.Application
	builder                   instructions.Builder
	instructionBuilder        instructions.InstructionBuilder
	assignmentBuilder         instructions.AssignmentBuilder
	valueBuilder              instructions.ValueBuilder
	applicationBuilder        applications.Builder
	attachmentBuilder         attachments.Builder
	attachmentVariableBuilder attachments.VariableBuilder
	parameterBuilder          parameters.Builder
	outputBuilder             instructions.OutputBuilder
}

func createApplication(
	grammarApp grammar_application.Application,
	criteriaApp application_criteria.Application,
	builder instructions.Builder,
	instructionBuilder instructions.InstructionBuilder,
	assignmentBuilder instructions.AssignmentBuilder,
	valueBuilder instructions.ValueBuilder,
	applicationBuilder applications.Builder,
	attachmentBuilder attachments.Builder,
	attachmentVariableBuilder attachments.VariableBuilder,
	parameterBuilder parameters.Builder,
	outputBuilder instructions.OutputBuilder,
) Application {
	out := application{
		grammarApp:                grammarApp,
		criteriaApp:               criteriaApp,
		builder:                   builder,
		instructionBuilder:        instructionBuilder,
		assignmentBuilder:         assignmentBuilder,
		valueBuilder:              valueBuilder,
		applicationBuilder:        applicationBuilder,
		attachmentBuilder:         attachmentBuilder,
		attachmentVariableBuilder: attachmentVariableBuilder,
		parameterBuilder:          parameterBuilder,
		outputBuilder:             outputBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(grammar grammars.Grammar, command commands.Command, script []byte) (instructions.Output, error) {
	tree, err := app.grammarApp.Execute(grammar, script)
	if err != nil {
		return nil, err
	}

	fmt.Printf("\n %s \n", tree.Bytes(true))

	index := uint(0)
	remaining := script
	instructionsList := []instructions.Instruction{}
	for {
		tree, err := app.grammarApp.Execute(grammar, remaining)
		if err != nil {
			break
		}

		retInstruction, err := app.instruction(
			index,
			grammar,
			command,
			tree,
		)

		if err != nil {
			break
		}

		remaining = tree.Remaining()
		instructionsList = append(instructionsList, retInstruction)
		index++

		if !tree.HasRemaining() {
			break
		}
	}

	instructions, err := app.builder.Create().WithList(instructionsList).Now()
	if err != nil {
		return nil, err
	}

	outputBuilder := app.outputBuilder.Create().WithInstructions(instructions)
	if len(remaining) > 0 {
		outputBuilder.WithRemaining(remaining)
	}

	return outputBuilder.Now()
}

func (app *application) instruction(
	index uint,
	grammar grammars.Grammar,
	command commands.Command,
	tree trees.Tree,
) (instructions.Instruction, error) {
	builder := app.instructionBuilder.Create()
	moduleCritria := command.ModuleDeclaration()
	moduleValue, err := app.criteriaApp.Execute(moduleCritria, tree)
	if err == nil {
		return builder.WithModule(string(moduleValue)).Now()
	}

	applicationDeclaration := command.ApplicationDeclaration()
	application, err := app.application(tree, applicationDeclaration)
	if err == nil {
		return builder.WithApplication(application).Now()
	}

	parameterDeclaration := command.ParameterDeclaration()
	parameter, err := app.parameter(tree, parameterDeclaration)
	if err == nil {
		return builder.WithParameter(parameter).Now()
	}

	variableAssignment := command.VariableAssignment()
	assignment, err := app.assignment(grammar, command, tree, variableAssignment)
	if err == nil {
		return builder.WithAssignment(assignment).Now()
	}

	attachment := command.Attachment()
	retAttachment, err := app.attachment(tree, attachment)
	if err == nil {
		return builder.WithAttachment(retAttachment).Now()
	}

	executionCriteria := command.Execution()
	executionValue, err := app.criteriaApp.Execute(executionCriteria, tree)
	if err == nil {
		return builder.WithExecution(string(executionValue)).Now()
	}

	str := fmt.Sprintf("the instruction (index: %d, value: %s) is invalid", index, string(tree.Bytes(true)))
	return nil, errors.New(str)
}

func (app *application) attachment(
	tree trees.Tree,
	attachmentCmd commands.Attachment,
) (attachments.Attachment, error) {
	applicationCriteria := attachmentCmd.Application()
	applicationValue, err := app.criteriaApp.Execute(applicationCriteria, tree)
	if err != nil {
		return nil, err
	}

	currentCriteria := attachmentCmd.Current()
	targetCriteria := attachmentCmd.Target()
	variable, err := app.attachmentVariable(tree, currentCriteria, targetCriteria)
	if err != nil {
		return nil, err
	}

	return app.attachmentBuilder.Create().
		WithVariable(variable).
		WithApplication(string(applicationValue)).
		Now()
}

func (app *application) attachmentVariable(
	tree trees.Tree,
	current criterias.Criteria,
	target criterias.Criteria,
) (attachments.Variable, error) {
	currentValue, err := app.criteriaApp.Execute(current, tree)
	if err != nil {
		return nil, err
	}

	targetValue, err := app.criteriaApp.Execute(target, tree)
	if err != nil {
		return nil, err
	}

	return app.attachmentVariableBuilder.Create().
		WithCurrent(string(currentValue)).
		WithTarget(string(targetValue)).
		Now()
}

func (app *application) assignment(
	grammar grammars.Grammar,
	command commands.Command,
	tree trees.Tree,
	assignmentCmd commands.VariableAssignment,
) (instructions.Assignment, error) {
	assigneeCriteria := assignmentCmd.Assignee()
	assigneeValue, err := app.criteriaApp.Execute(assigneeCriteria, tree)
	if err != nil {
		return nil, err
	}

	valueCmd := assignmentCmd.Value()
	value, err := app.value(grammar, command, tree, valueCmd)
	if err != nil {
		return nil, err
	}

	return app.assignmentBuilder.Create().
		WithVariable(string(assigneeValue)).
		WithValue(value).
		Now()
}

func (app *application) value(
	grammar grammars.Grammar,
	command commands.Command,
	tree trees.Tree,
	valueCmd commands.Value,
) (instructions.Value, error) {
	builder := app.valueBuilder.Create()
	variableCriteria := valueCmd.Variable()
	variableValue, err := app.criteriaApp.Execute(variableCriteria, tree)
	if err == nil {
		builder.WithInput(string(variableValue))
	}

	constantCriteria := valueCmd.Constant()
	constantValue, err := app.criteriaApp.Execute(constantCriteria, tree)
	if err == nil {
		builder.WithString(string(constantValue))
	}

	instructionsCriteria := valueCmd.Instructions()
	instructionsValue, err := app.criteriaApp.Execute(instructionsCriteria, tree)

	fmt.Printf("-> %s\n", tree.Bytes(true))

	if err == nil {
		subOutput, err := app.Execute(grammar, command, instructionsValue)
		if err != nil {
			return nil, err
		}

		if subOutput.HasRemaining() {
			str := fmt.Sprintf("the instruction's value were NOT expected to contain remaining data")
			return nil, errors.New(str)
		}

		subInstructions := subOutput.Instructions()
		builder.WithInstructions(subInstructions)
	}

	executionCriteria := valueCmd.Execution()
	executionValue, err := app.criteriaApp.Execute(executionCriteria, tree)
	if err == nil {
		builder.WithExecution(string(executionValue))
	}

	return builder.Now()
}

func (app *application) parameter(
	tree trees.Tree,
	parameterDeclaration commands.ParameterDeclaration,
) (parameters.Parameter, error) {
	inputCriteria := parameterDeclaration.Input()
	inputName, err := app.criteriaApp.Execute(inputCriteria, tree)
	if err == nil {
		nameStr := string(inputName)
		return app.parameterBuilder.Create().WithName(nameStr).IsInput().Now()
	}

	outputCriteria := parameterDeclaration.Output()
	outputName, err := app.criteriaApp.Execute(outputCriteria, tree)
	if err != nil {
		return nil, err
	}

	nameStr := string(outputName)
	return app.parameterBuilder.Create().WithName(nameStr).Now()
}

func (app *application) application(
	tree trees.Tree,
	applicationDeclaration commands.ApplicationDeclaration,
) (applications.Application, error) {
	moduleCriteria := applicationDeclaration.Module()
	moduleName, err := app.criteriaApp.Execute(moduleCriteria, tree)
	if err != nil {
		return nil, err
	}

	nameCriteria := applicationDeclaration.Name()
	name, err := app.criteriaApp.Execute(nameCriteria, tree)
	if err != nil {
		return nil, err
	}

	nameStr := string(name)
	moduleNameStr := string(moduleName)
	return app.applicationBuilder.Create().
		WithModule(moduleNameStr).
		WithName(nameStr).
		Now()
}
