package instructions

import (
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
	return nil, nil
}

func (app *application) value(
	tree trees.Tree,
	criteria criterias.Criteria,
) (instructions.Value, error) {
	/*value, err := app.criteriaApp.Execute(criteria, tree)
	if err != nil {
		return nil, nil
	}

	valueStr := string(value)
	valueBuilder := app.valueBuilder.Create()
	if assignmentIns, ok := inVariables[valueStr]; ok {
		variableIns := assignmentIns.Value()
		if variableIns.IsInput() {
			input := variableIns.Input()
			valueBuilder.WithInput(input)
		}

		if variableIns.IsString() {
			str := variableIns.String()
			valueBuilder.WithString(str)
		}

		if variableIns.IsExecution() {
			execution := variableIns.Execution()
			valueBuilder.WithExecution(execution)
		}
	} else if isInput, ok := inParameters[valueStr]; ok {
		if !isInput {
			str := fmt.Sprintf("the output parameter (name: %s) cannot be used as a value in an assignment", valueStr)
			return nil, errors.New(str)
		}

		valueBuilder.WithInput(valueStr)
	} else if appIns, ok := inApplications[valueStr]; ok {
		valueBuilder.WithExecution(appIns)
	} else {
		valueBuilder.WithString(valueStr)
	}

	return valueBuilder.Now()*/
	return nil, nil
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
		return nil, nil
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
		return nil, nil
	}

	nameCriteria := applicationDeclaration.Name()
	name, err := app.criteriaApp.Execute(nameCriteria, tree)
	if err != nil {
		return nil, nil
	}

	nameStr := string(name)
	moduleNameStr := string(moduleName)
	return app.applicationBuilder.Create().
		WithModule(moduleNameStr).
		WithName(nameStr).
		Now()
}
