package programs

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/domain/instructions"
	instruction_applications "github.com/steve-care-software/webx/domain/instructions/applications"
	"github.com/steve-care-software/webx/domain/instructions/attachments"
	"github.com/steve-care-software/webx/domain/instructions/parameters"
	"github.com/steve-care-software/webx/domain/programs"
	"github.com/steve-care-software/webx/domain/programs/modules"
)

type application struct {
	builder            programs.Builder
	instructionBuilder programs.InstructionBuilder
	applicationBuilder programs.ApplicationBuilder
	attachmentsBuilder programs.AttachmentsBuilder
	attachmentBuilder  programs.AttachmentBuilder
	assignmentBuilder  programs.AssignmentBuilder
	valueBuilder       programs.ValueBuilder
	modules            modules.Modules
}

func createApplication(
	builder programs.Builder,
	instructionBuilder programs.InstructionBuilder,
	applicationBuilder programs.ApplicationBuilder,
	attachmentsBuilder programs.AttachmentsBuilder,
	attachmentBuilder programs.AttachmentBuilder,
	assignmentBuilder programs.AssignmentBuilder,
	valueBuilder programs.ValueBuilder,
) Application {
	return createApplicationInternally(
		builder,
		instructionBuilder,
		applicationBuilder,
		attachmentsBuilder,
		attachmentBuilder,
		assignmentBuilder,
		valueBuilder,
		nil,
	)
}

func createApplicationWithModules(
	builder programs.Builder,
	instructionBuilder programs.InstructionBuilder,
	applicationBuilder programs.ApplicationBuilder,
	attachmentsBuilder programs.AttachmentsBuilder,
	attachmentBuilder programs.AttachmentBuilder,
	assignmentBuilder programs.AssignmentBuilder,
	valueBuilder programs.ValueBuilder,
	modules modules.Modules,
) Application {
	return createApplicationInternally(
		builder,
		instructionBuilder,
		applicationBuilder,
		attachmentsBuilder,
		attachmentBuilder,
		assignmentBuilder,
		valueBuilder,
		modules,
	)
}

func createApplicationInternally(
	builder programs.Builder,
	instructionBuilder programs.InstructionBuilder,
	applicationBuilder programs.ApplicationBuilder,
	attachmentsBuilder programs.AttachmentsBuilder,
	attachmentBuilder programs.AttachmentBuilder,
	assignmentBuilder programs.AssignmentBuilder,
	valueBuilder programs.ValueBuilder,
	modules modules.Modules,
) Application {
	out := application{
		builder:            builder,
		instructionBuilder: instructionBuilder,
		applicationBuilder: applicationBuilder,
		attachmentsBuilder: attachmentsBuilder,
		attachmentBuilder:  attachmentBuilder,
		assignmentBuilder:  assignmentBuilder,
		valueBuilder:       valueBuilder,
		modules:            modules,
	}

	return &out
}

// Execute executes instructions to build a program instance
func (app *application) Execute(instructions instructions.Instructions) (programs.Program, error) {
	list := instructions.List()
	inModules := map[string]modules.Module{}
	inApplications := map[string]programs.Application{}
	inParameters := map[string]bool{}
	inAssignments := map[string]programs.Assignment{}
	inInstructions := []programs.Instruction{}
	for idx, oneInstruction := range list {
		outModules, outApplications, outParameters, outAssignments, outInstructions, err := app.instruction(
			uint(idx),
			oneInstruction,
			inModules,
			inApplications,
			inParameters,
			inAssignments,
			inInstructions,
		)

		if err != nil {
			return nil, err
		}

		inModules = outModules
		inApplications = outApplications
		inParameters = outParameters
		inAssignments = outAssignments
		inInstructions = outInstructions
	}

	outputs := []string{}
	for name, isInput := range inParameters {
		if isInput {
			continue
		}

		outputs = append(outputs, name)
	}

	builder := app.builder.Create().WithInstructions(inInstructions)
	if len(outputs) > 0 {
		builder.WithOutputs(outputs)
	}

	return builder.Now()
}

func (app *application) instruction(
	index uint,
	instruction instructions.Instruction,
	inModules map[string]modules.Module,
	inApplications map[string]programs.Application,
	inParameters map[string]bool,
	inAssignments map[string]programs.Assignment,
	inInstructions []programs.Instruction,
) (map[string]modules.Module, map[string]programs.Application, map[string]bool, map[string]programs.Assignment, []programs.Instruction, error) {
	if instruction.IsModule() {
		name := instruction.Module()
		outModules, err := app.module(name, inModules)
		if err != nil {
			return nil, nil, nil, nil, nil, err
		}

		return outModules, inApplications, inParameters, inAssignments, inInstructions, nil
	}

	if instruction.IsApplication() {
		insApp := instruction.Application()
		outApplications, err := app.application(insApp, inModules, inApplications)
		if err != nil {
			return nil, nil, nil, nil, nil, err
		}

		return inModules, outApplications, inParameters, inAssignments, inInstructions, nil
	}

	if instruction.IsParameter() {
		insParameter := instruction.Parameter()
		outParameters, err := app.parameter(insParameter, inParameters)
		if err != nil {
			return nil, nil, nil, nil, nil, err
		}

		return inModules, inApplications, outParameters, inAssignments, inInstructions, nil
	}

	if instruction.IsAssignment() {
		assignment := instruction.Assignment()
		outAssignments, newAssignment, err := app.assignment(index, assignment, inParameters, inAssignments, inApplications)
		if err != nil {
			return nil, nil, nil, nil, nil, err
		}

		instruction, err := app.instructionBuilder.Create().WithAssignment(newAssignment).Now()
		if err != nil {
			return nil, nil, nil, nil, nil, err
		}

		inInstructions = append(inInstructions, instruction)
		return inModules, inApplications, inParameters, outAssignments, inInstructions, nil
	}

	if instruction.IsAttachment() {
		attachment := instruction.Attachment()
		outApplications, err := app.attachment(attachment, inParameters, inAssignments, inApplications)
		if err != nil {
			return nil, nil, nil, nil, nil, err
		}

		return inModules, outApplications, inParameters, inAssignments, inInstructions, nil
	}

	execution := instruction.Execution()
	instructionBuilder := app.instructionBuilder.Create()
	if executedApp, ok := inApplications[execution]; ok {
		instructionBuilder.WithExecution(executedApp)
	} else {
		str := fmt.Sprintf("the application's execution (index: %d, name: %s) is invalid because the application is undefined", index, execution)
		return nil, nil, nil, nil, nil, errors.New(str)
	}

	ins, err := instructionBuilder.Now()
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	inInstructions = append(inInstructions, ins)
	return inModules, inApplications, inParameters, inAssignments, inInstructions, nil
}

func (app *application) attachment(
	attachment attachments.Attachment,
	inParameters map[string]bool,
	inAssignments map[string]programs.Assignment,
	inApplications map[string]programs.Application,
) (map[string]programs.Application, error) {
	variable := attachment.Variable()
	currentValue, err := app.attachmentValue(variable, inParameters, inAssignments)
	if err != nil {
		return nil, err
	}

	applicationName := attachment.Application()
	if appIns, ok := inApplications[applicationName]; ok {
		target := variable.Target()
		attachment, err := app.attachmentBuilder.Create().WithValue(currentValue).WithLocal(target).Now()
		if err != nil {
			return nil, err
		}

		attachmentsList := []programs.Attachment{}
		if appIns.HasAttachments() {
			attachmentsList = appIns.Attachments().List()
		}

		attachmentsList = append(attachmentsList, attachment)
		attachments, err := app.attachmentsBuilder.Create().WithList(attachmentsList).Now()
		if err != nil {
			return nil, err
		}

		name := appIns.Name()
		module := appIns.Module()
		updatedAppIns, err := app.applicationBuilder.Create().WithName(name).WithModule(module).WithAttachments(attachments).Now()
		if err != nil {
			return nil, err
		}

		inApplications[name] = updatedAppIns
		return inApplications, nil
	}

	str := fmt.Sprintf("the application (name: %s) is undeclared and therefore cannot be used in an attachment", applicationName)
	return nil, errors.New(str)
}

func (app *application) attachmentValue(
	variable attachments.Variable,
	inParameters map[string]bool,
	inAssignments map[string]programs.Assignment,
) (programs.Value, error) {
	current := variable.Current()
	if currentIns, ok := inAssignments[current]; ok {
		return currentIns.Value(), nil
	}

	if isInput, ok := inParameters[current]; ok {
		if !isInput {
			str := fmt.Sprintf("the output variable (name: %s) cannot be used in attachment", current)
			return nil, errors.New(str)
		}

		return app.valueBuilder.Create().WithInput(current).Now()
	}

	str := fmt.Sprintf("the current variable (name: %s) is undeclared and therefore cannot be used in an attachment", current)
	return nil, errors.New(str)
}

func (app *application) assignment(
	index uint,
	assignment instructions.Assignment,
	inParameters map[string]bool,
	inAssignments map[string]programs.Assignment,
	inApplications map[string]programs.Application,
) (map[string]programs.Assignment, programs.Assignment, error) {
	valueIns, err := app.value(index, assignment, inParameters, inAssignments, inApplications)
	if err != nil {
		return nil, nil, err
	}

	variable := assignment.Variable()
	assignmentIns, err := app.assignmentBuilder.Create().WithIndex(index).WithName(variable).WithValue(valueIns).Now()
	if err != nil {
		return nil, nil, err
	}

	inAssignments[variable] = assignmentIns
	return inAssignments, assignmentIns, nil
}

func (app *application) value(
	index uint,
	assignment instructions.Assignment,
	inParameters map[string]bool,
	inAssignments map[string]programs.Assignment,
	inApplications map[string]programs.Application,
) (programs.Value, error) {
	variable := assignment.Variable()
	value := assignment.Value()
	builder := app.valueBuilder.Create()
	if value.IsInput() {
		input := value.Input()
		if isInput, ok := inParameters[input]; ok {
			if !isInput {
				str := fmt.Sprintf("the assignment (index: %d, variable: %s) is using an output variable (name: %s) as value", index, variable, input)
				return nil, errors.New(str)
			}

			builder.WithInput(input)
		} else {
			str := fmt.Sprintf("the assignment (index: %d, variable: %s) is using an undefined input variable (name: %s) as value", index, variable, input)
			return nil, errors.New(str)
		}
	}

	if value.IsString() {
		str := value.String()
		if variable, ok := inAssignments[str]; ok {
			return variable.Value(), nil
		}

		builder.WithString(str)
	}

	if value.IsInstructions() {
		subInstructions := value.Instructions()
		subProgram, err := app.Execute(subInstructions)
		if err != nil {
			return nil, err
		}

		builder.WithProgram(subProgram)
	}

	if value.IsExecution() {
		execution := value.Execution()
		if executedApp, ok := inApplications[execution]; ok {
			builder.WithExecution(executedApp)
		} else {
			str := fmt.Sprintf("the assignment (index: %d, variable: %s) is using an undefined application execution (name: %s) as value", index, variable, execution)
			return nil, errors.New(str)
		}
	}

	return builder.Now()
}

func (app *application) parameter(
	parameter parameters.Parameter,
	inParameters map[string]bool,
) (map[string]bool, error) {
	name := parameter.Name()
	if _, ok := inParameters[name]; ok {
		str := fmt.Sprintf("the parameter (name: %s, isInput: %t) is already declared", name, parameter.IsInput())
		return nil, errors.New(str)
	}

	inParameters[name] = parameter.IsInput()
	return inParameters, nil
}

func (app *application) application(
	application instruction_applications.Application,
	inModules map[string]modules.Module,
	inApplications map[string]programs.Application,
) (map[string]programs.Application, error) {
	name := application.Name()
	if _, ok := inApplications[name]; ok {
		str := fmt.Sprintf("the application (name: %s) is already declared", name)
		return nil, errors.New(str)
	}

	module := application.Module()
	if _, ok := inModules[module]; !ok {
		str := fmt.Sprintf("the module (name: %s) is undefined but used in the application declaration (name: %s)", module, name)
		return nil, errors.New(str)
	}

	ins, err := app.applicationBuilder.Create().WithName(name).WithModule(inModules[module]).Now()
	if err != nil {
		return nil, err
	}

	inApplications[name] = ins
	return inApplications, nil
}

func (app *application) module(
	name string,
	modules map[string]modules.Module,
) (map[string]modules.Module, error) {
	if app.modules == nil {
		str := fmt.Sprintf("the module (name: %s) is undefined", name)
		return nil, errors.New(str)
	}

	module, err := app.modules.Fetch(name)
	if err != nil {
		return nil, err
	}

	if _, ok := modules[name]; ok {
		str := fmt.Sprintf("the module (name: %s) is already loaded", name)
		return nil, errors.New(str)
	}

	modules[name] = module
	return modules, nil
}
