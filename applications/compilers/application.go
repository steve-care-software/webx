package compilers

import (
	"github.com/steve-care-software/webx/applications/creates"
	grammar_application "github.com/steve-care-software/webx/applications/grammars"
	interpreter_application "github.com/steve-care-software/webx/applications/interpreters"
	"github.com/steve-care-software/webx/domain/compilers"
	"github.com/steve-care-software/webx/domain/instructions"
	"github.com/steve-care-software/webx/domain/instructions/applications"
	"github.com/steve-care-software/webx/domain/instructions/attachments"
)

type application struct {
	grammarApp grammar_application.Application
	//criteriaApp               criteria_application.Application
	interpreterApp            interpreter_application.Application
	createApp                 creates.Application
	instructionsBuilder       instructions.Builder
	instructionBuilder        instructions.InstructionBuilder
	applicationBuilder        applications.Builder
	attachmentBuilder         attachments.Builder
	attachmentVariableBuilder attachments.VariableBuilder
	assignmentBuilder         instructions.AssignmentBuilder
	valueBuilder              instructions.ValueBuilder
	outputBuilder             instructions.OutputBuilder
}

func createApplication(
	grammarApp grammar_application.Application,
	//criteriaApp criteria_application.Application,
	interpreterApp interpreter_application.Application,
	createApp creates.Application,
	instructionsBuilder instructions.Builder,
	instructionBuilder instructions.InstructionBuilder,
	applicationBuilder applications.Builder,
	attachmentBuilder attachments.Builder,
	attachmentVariableBuilder attachments.VariableBuilder,
	assignmentBuilder instructions.AssignmentBuilder,
	valueBuilder instructions.ValueBuilder,
	outputBuilder instructions.OutputBuilder,
) Application {
	out := application{
		grammarApp: grammarApp,
		//criteriaApp:               criteriaApp,
		interpreterApp:            interpreterApp,
		createApp:                 createApp,
		instructionsBuilder:       instructionsBuilder,
		instructionBuilder:        instructionBuilder,
		applicationBuilder:        applicationBuilder,
		attachmentBuilder:         attachmentBuilder,
		attachmentVariableBuilder: attachmentVariableBuilder,
		assignmentBuilder:         assignmentBuilder,
		valueBuilder:              valueBuilder,
		outputBuilder:             outputBuilder,
	}

	return &out
}

// Execute eecutes a compiler
func (app *application) Execute(compiler compilers.Compiler, script []byte) (instructions.Output, error) {
	/*remaining := script
	elementsList := compiler.Elements().List()
	application, err := app.applicationBuilder.Create().WithModule(ExecuteInstructionsModuleName).WithName("executeInstructionsApp").Now()
	if err != nil {
		return nil, err
	}

	applicationDeclarationIns, err := app.instructionBuilder.Create().WithApplication(application).Now()
	if err != nil {
		return nil, err
	}

	instructionsList := []instructions.Instruction{
		applicationDeclarationIns,
	}

	applicationName := application.Name()
	for elementIndex, oneElement := range elementsList {
		grammar := oneElement.Grammar()
		tree, err := app.grammarApp.Execute(grammar, remaining)
		if err != nil {
			return nil, err
		}

		codeVariable := fmt.Sprintf("ins%d", elementIndex)
		parameters := oneElement.Parameters().List()
		for parameterIndex, oneParameter := range parameters {
			parameterValue := oneParameter.Value()
			valueBuilder := app.valueBuilder.Create()
			if parameterValue.IsConstant() {
				constant := parameterValue.Constant()
				valueBuilder.WithString(constant)
			}

			if parameterValue.IsCriteria() {
				criteria := parameterValue.Criteria()
				found, err := app.criteriaApp.Execute(criteria, tree)
				if err != nil {
					return nil, err
				}

				valueBuilder.WithString(string(found))
				if err != nil {
					return nil, err
				}
			}

			value, err := valueBuilder.Now()
			if err != nil {
				return nil, err
			}

			currentName := fmt.Sprintf("%s%d", codeVariable, parameterIndex)
			assignment, err := app.assignmentBuilder.Create().WithVariable(currentName).WithValue(value).Now()
			if err != nil {
				return nil, err
			}

			assignmentIns, err := app.instructionBuilder.Create().WithAssignment(assignment).Now()
			if err != nil {
				return nil, err
			}

			name := oneParameter.Name()
			attachmentVariable, err := app.attachmentVariableBuilder.Create().WithCurrent(currentName).WithTarget(name).Now()
			if err != nil {
				return nil, err
			}

			attachment, err := app.attachmentBuilder.Create().WithVariable(attachmentVariable).WithApplication(applicationName).Now()
			if err != nil {
				return nil, err
			}

			attachmentIns, err := app.instructionBuilder.Create().WithAttachment(attachment).Now()
			if err != nil {
				return nil, err
			}

			instructionsList = append(instructionsList, assignmentIns)
			instructionsList = append(instructionsList, attachmentIns)
		}

		execution := oneElement.Execution()
		instructions := execution.Instructions()
		codeValue, err := app.valueBuilder.Create().WithInstructions(instructions).Now()
		if err != nil {
			return nil, err
		}

		codeAssignment, err := app.assignmentBuilder.Create().WithVariable(codeVariable).WithValue(codeValue).Now()
		if err != nil {
			return nil, err
		}

		codeAssignmentIns, err := app.instructionBuilder.Create().WithAssignment(codeAssignment).Now()
		if err != nil {
			return nil, err
		}

		target := execution.Parameter()
		codeAttachmentVariable, err := app.attachmentVariableBuilder.Create().WithCurrent(codeVariable).WithTarget(target).Now()
		if err != nil {
			return nil, err
		}

		codeAttachment, err := app.attachmentBuilder.Create().WithVariable(codeAttachmentVariable).WithApplication(applicationName).Now()
		if err != nil {
			return nil, err
		}

		codeAttachmentIns, err := app.instructionBuilder.Create().WithAttachment(codeAttachment).Now()
		if err != nil {
			return nil, err
		}

		instructionsList = append(instructionsList, codeAssignmentIns)
		instructionsList = append(instructionsList, codeAttachmentIns)
		remaining = tree.Remaining()
	}

	instructions, err := app.instructionsBuilder.Create().WithList(instructionsList).Now()
	if err != nil {
		return nil, err
	}

	builder := app.outputBuilder.Create().WithInstructions(instructions)
	if remaining != nil {
		builder.WithRemaining(remaining)
	}

	return builder.Now()*/
	return nil, nil
}
