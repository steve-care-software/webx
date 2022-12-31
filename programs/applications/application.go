package applications

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
	"github.com/steve-care-software/webx/programs/domain/instructions"
	instructions_application "github.com/steve-care-software/webx/programs/domain/instructions/applications"
	"github.com/steve-care-software/webx/programs/domain/instructions/attachments"
	instructions_module "github.com/steve-care-software/webx/programs/domain/instructions/modules"
	"github.com/steve-care-software/webx/programs/domain/instructions/parameters"
	"github.com/steve-care-software/webx/programs/domain/programs"
	"github.com/steve-care-software/webx/programs/domain/programs/modules"
)

type application struct {
	builder             programs.Builder
	instructionsBuilder programs.InstructionsBuilder
	instructionBuilder  programs.InstructionBuilder
	applicationBuilder  programs.ApplicationBuilder
	attachmentsBuilder  programs.AttachmentsBuilder
	attachmentBuilder   programs.AttachmentBuilder
	valueBuilder        programs.ValueBuilder
	hashAdapter         hash.Adapter
	nameBytesToStringFn NameBytesToString
}

func createApplication(
	builder programs.Builder,
	instructionsBuilder programs.InstructionsBuilder,
	instructionBuilder programs.InstructionBuilder,
	applicationBuilder programs.ApplicationBuilder,
	attachmentsBuilder programs.AttachmentsBuilder,
	attachmentBuilder programs.AttachmentBuilder,
	valueBuilder programs.ValueBuilder,
	hashAdapter hash.Adapter,
	nameBytesToStringFn NameBytesToString,
) Application {
	out := application{
		builder:             builder,
		instructionsBuilder: instructionsBuilder,
		instructionBuilder:  instructionBuilder,
		applicationBuilder:  applicationBuilder,
		attachmentsBuilder:  attachmentsBuilder,
		attachmentBuilder:   attachmentBuilder,
		valueBuilder:        valueBuilder,
		hashAdapter:         hashAdapter,
		nameBytesToStringFn: nameBytesToStringFn,
	}
	return &out
}

// Compile compiles modules and instructions to a program instance
func (app *application) Compile(modulesIns modules.Modules, instructions instructions.Instructions) (programs.Program, error) {
	list := instructions.List()
	inModules := map[string]modules.Module{}
	inApplications := map[string]programs.Application{}
	inParameters := map[string]*parameter{}
	inValues := map[string]programs.Value{}
	inInstructions := []programs.Instruction{}
	for idx, oneInstruction := range list {
		outModules, outApplications, outParameters, outValues, outInstructions, err := app.compileInstruction(
			oneInstruction,
			inModules,
			inApplications,
			inParameters,
			inValues,
			inInstructions,
			modulesIns,
		)

		if err != nil {
			str := fmt.Sprintf("there was an error at instruction (index: %d): %s", idx, err.Error())
			return nil, errors.New(str)
		}

		inModules = outModules
		inApplications = outApplications
		inParameters = outParameters
		inValues = outValues
		inInstructions = outInstructions
	}

	outputs := []uint{}
	for _, parameter := range inParameters {
		if parameter.parameter.IsInput() {
			continue
		}

		outputs = append(outputs, parameter.index)
	}

	ins, err := app.instructionsBuilder.Create().WithList(inInstructions).Now()
	if err != nil {
		return nil, err
	}

	builder := app.builder.Create().WithInstructions(ins)
	if len(outputs) > 0 {
		builder.WithOutputs(outputs)
	}

	return builder.Now()
}

func (app *application) compileInstruction(
	instruction instructions.Instruction,
	inModules map[string]modules.Module,
	inApplications map[string]programs.Application,
	inParameters map[string]*parameter,
	inValues map[string]programs.Value,
	inInstructions []programs.Instruction,
	allModules modules.Modules,
) (map[string]modules.Module, map[string]programs.Application, map[string]*parameter, map[string]programs.Value, []programs.Instruction, error) {
	if instruction.IsModule() {
		name := instruction.Module()
		outModules, err := app.compileModule(name, inModules, allModules)
		if err != nil {
			return nil, nil, nil, nil, nil, err
		}

		return outModules, inApplications, inParameters, inValues, inInstructions, nil
	}

	if instruction.IsApplication() {
		insApp := instruction.Application()
		outApplications, err := app.compileApplication(insApp, inModules, inApplications)
		if err != nil {
			return nil, nil, nil, nil, nil, err
		}

		return inModules, outApplications, inParameters, inValues, inInstructions, nil
	}

	if instruction.IsParameter() {
		insParameter := instruction.Parameter()
		outParameters, err := app.compileParameter(insParameter, inParameters)
		if err != nil {
			return nil, nil, nil, nil, nil, err
		}

		return inModules, inApplications, outParameters, inValues, inInstructions, nil
	}

	if instruction.IsAssignment() {
		assignment := instruction.Assignment()
		valueIns, err := app.compileValue(assignment, inParameters, inApplications, allModules)
		if err != nil {
			return nil, nil, nil, nil, nil, err
		}

		variableName := assignment.Variable()
		variableNameStr := app.nameBytesToStringFn(variableName)

		outValues := inValues
		outValues[variableNameStr] = valueIns

		return inModules, inApplications, inParameters, outValues, inInstructions, nil
	}

	if instruction.IsAttachment() {
		attachment := instruction.Attachment()
		outApplications, err := app.compileAttachment(attachment, inParameters, inValues, inApplications, allModules)
		if err != nil {
			return nil, nil, nil, nil, nil, err
		}

		return inModules, outApplications, inParameters, inValues, inInstructions, nil
	}

	execution := instruction.Execution()
	executionNameStr := app.nameBytesToStringFn(execution)
	instructionBuilder := app.instructionBuilder.Create()
	if executedApp, ok := inApplications[executionNameStr]; ok {
		instructionBuilder.WithExecution(executedApp)
	} else {
		str := fmt.Sprintf("the application's execution (name: %s) is invalid because the application is undefined", executionNameStr)
		return nil, nil, nil, nil, nil, errors.New(str)
	}

	ins, err := instructionBuilder.Now()
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	inInstructions = append(inInstructions, ins)
	return inModules, inApplications, inParameters, inValues, inInstructions, nil
}

func (app *application) compileAttachment(
	attachment attachments.Attachment,
	inParameters map[string]*parameter,
	inValues map[string]programs.Value,
	inApplications map[string]programs.Application,
	allModules modules.Modules,
) (map[string]programs.Application, error) {
	variable := attachment.Variable()
	currentValue, err := app.compileAttachmentValue(variable, inParameters, inValues, allModules)
	if err != nil {
		return nil, err
	}

	applicationName := attachment.Application()
	applicationNameStr := app.nameBytesToStringFn(applicationName)
	if appIns, ok := inApplications[applicationNameStr]; ok {
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

		index := appIns.Index()
		module := appIns.Module()
		updatedAppIns, err := app.applicationBuilder.Create().WithIndex(index).WithModule(module).WithAttachments(attachments).Now()
		if err != nil {
			return nil, err
		}

		inApplications[applicationNameStr] = updatedAppIns
		return inApplications, nil
	}

	str := fmt.Sprintf("the application (name: %s) is undeclared and therefore cannot be used in an attachment", applicationName)
	return nil, errors.New(str)
}

func (app *application) compileAttachmentValue(
	variable attachments.Variable,
	inParameters map[string]*parameter,
	inValues map[string]programs.Value,
	allModules modules.Modules,
) (programs.Value, error) {
	current := variable.Current()
	currentNameStr := app.nameBytesToStringFn(current)
	if currentIns, ok := inValues[currentNameStr]; ok {
		return currentIns, nil
	}

	if parameter, ok := inParameters[currentNameStr]; ok {
		if !parameter.parameter.IsInput() {
			str := fmt.Sprintf("the output variable (name: %s, index: %d) cannot be used in attachment", currentNameStr, parameter.index)
			return nil, errors.New(str)
		}

		return app.valueBuilder.Create().WithInput(parameter.index).Now()
	}

	str := fmt.Sprintf("the current variable (name: %s) is undeclared and therefore cannot be used in an attachment", currentNameStr)
	return nil, errors.New(str)
}

func (app *application) compileValue(
	assignment instructions.Assignment,
	inParameters map[string]*parameter,
	inApplications map[string]programs.Application,
	allModules modules.Modules,
) (programs.Value, error) {
	variable := assignment.Variable()
	variableNameStr := app.nameBytesToStringFn(variable)

	value := assignment.Value()
	builder := app.valueBuilder.Create()
	if value.IsVariable() {
		assignedVariable := value.Variable()
		assignedVariableNameStr := app.nameBytesToStringFn(assignedVariable)
		if parameter, ok := inParameters[assignedVariableNameStr]; ok {
			if !parameter.parameter.IsInput() {
				str := fmt.Sprintf("the assignment (name: %s) is using an output variable (nme: %s) as value", variableNameStr, assignedVariableNameStr)
				return nil, errors.New(str)
			}

			builder.WithInput(parameter.index)
		} else {
			str := fmt.Sprintf("the assignment (name: %s) is using an undefined parameter (name: %s) as value", variableNameStr, assignedVariableNameStr)
			return nil, errors.New(str)
		}
	}

	if value.IsConstant() {
		constant := value.Constant()
		builder.WithConstant(constant)
	}

	if value.IsInstructions() {
		subInstructions := value.Instructions()
		subProgram, err := app.Compile(allModules, subInstructions)
		if err != nil {
			return nil, err
		}

		builder.WithProgram(subProgram)
	}

	if value.IsExecution() {
		execution := value.Execution()
		executionNameStr := app.nameBytesToStringFn(execution)
		if executedApp, ok := inApplications[executionNameStr]; ok {
			builder.WithExecution(executedApp)
		} else {
			str := fmt.Sprintf("the assignment (name: %s) is using an undefined application execution (name: %s) as value", variableNameStr, executionNameStr)
			return nil, errors.New(str)
		}
	}

	return builder.Now()
}

func (app *application) compileParameter(
	parameterIns parameters.Parameter,
	inParameters map[string]*parameter,
) (map[string]*parameter, error) {
	name := parameterIns.Name()
	parameterNameStr := app.nameBytesToStringFn(name)
	if _, ok := inParameters[parameterNameStr]; ok {
		str := fmt.Sprintf("the parameter (name: %s, isInput: %t) is already declared", parameterNameStr, parameterIns.IsInput())
		return nil, errors.New(str)
	}

	inParameters[parameterNameStr] = &parameter{
		index:     uint(len(inParameters)),
		parameter: parameterIns,
	}

	return inParameters, nil
}

func (app *application) compileApplication(
	application instructions_application.Application,
	inModules map[string]modules.Module,
	inApplications map[string]programs.Application,
) (map[string]programs.Application, error) {
	name := application.Name()
	appNameStr := app.nameBytesToStringFn(name)
	if _, ok := inApplications[appNameStr]; ok {
		str := fmt.Sprintf("the application (name: %s) is already declared", appNameStr)
		return nil, errors.New(str)
	}

	module := application.Module()
	moduleNameStr := app.nameBytesToStringFn(module)
	if _, ok := inModules[moduleNameStr]; !ok {
		str := fmt.Sprintf("the module (name: %string) is undefined but used in the application declaration (name: %s)", moduleNameStr, appNameStr)
		return nil, errors.New(str)
	}

	appIndex := uint(len(inApplications))
	ins, err := app.applicationBuilder.Create().WithIndex(appIndex).WithModule(inModules[moduleNameStr]).Now()
	if err != nil {
		return nil, err
	}

	inApplications[appNameStr] = ins
	return inApplications, nil
}

func (app *application) compileModule(
	insModule instructions_module.Module,
	loadedModules map[string]modules.Module,
	allModules modules.Modules,
) (map[string]modules.Module, error) {
	index := insModule.Index()
	module, err := allModules.Fetch(index)
	if err != nil {
		return nil, err
	}

	name := insModule.Name()
	moduleNameStr := app.nameBytesToStringFn(name)
	if _, ok := loadedModules[moduleNameStr]; ok {
		str := fmt.Sprintf("the module (index: %d, name: %s) is already loaded", index, moduleNameStr)
		return nil, errors.New(str)
	}

	loadedModules[moduleNameStr] = module
	return loadedModules, nil
}

// Execute executes a program
func (app *application) Execute(input map[uint]interface{}, program programs.Program) (map[uint]interface{}, error) {
	valueHashes := map[string]interface{}{}
	valueIndexes := map[uint]interface{}{}
	instructions := program.Instructions().List()
	for idx, oneInstruction := range instructions {
		if oneInstruction.IsValue() {
			value := oneInstruction.Value()
			valueHash := value.Hash().String()

			ins, err := app.executeValue(input, valueHashes, value)
			if err != nil {
				str := fmt.Sprintf("there was an error while executing an assignment (index: %d. value's hash: %s): %s", idx, valueHash, err.Error())
				return nil, errors.New(str)
			}

			valueHashes[valueHash] = ins
			valueIndexes[uint(idx)] = ins
			continue
		}

		execution := oneInstruction.Execution()
		_, err := app.execute(input, valueHashes, execution)
		if err != nil {
			hash := execution.Hash().String()
			index := execution.Index()
			str := fmt.Sprintf("there was an error while executing an application (index: %d, application's hash: %s, application's index: %d): %s", idx, hash, index, err.Error())
			return nil, errors.New(str)
		}
	}

	filtered := map[uint]interface{}{}
	if program.HasOutputs() {
		outputs := program.Outputs()
		for _, oneOutput := range outputs {
			if ins, ok := valueIndexes[oneOutput]; ok {
				filtered[oneOutput] = ins
				continue
			}

			str := fmt.Sprintf("the program has an output parameter (%d), but the executed program does not contain that value", oneOutput)
			return nil, errors.New(str)
		}
	}

	return filtered, nil
}

func (app *application) executeValue(input map[uint]interface{}, values map[string]interface{}, value programs.Value) (interface{}, error) {
	content := value.Content()
	if content.IsInput() {
		pInputIndex := content.Input()
		if ins, ok := input[*pInputIndex]; ok {
			return ins, nil
		}

		str := fmt.Sprintf("the requested input variable (index: %d) is undefined", *pInputIndex)
		return nil, errors.New(str)
	}

	if content.IsConstant() {
		return content.Constant(), nil
	}

	if content.IsProgram() {
		subProgram := content.Program()
		subProgramOutput, err := app.Execute(input, subProgram)
		if err != nil {
			return nil, err
		}

		return subProgramOutput, nil
	}

	execution := content.Execution()
	return app.execute(input, values, execution)
}

func (app *application) execute(input map[uint]interface{}, values map[string]interface{}, execution programs.Application) (interface{}, error) {
	module := execution.Module()
	parameters := map[uint]interface{}{}
	if execution.HasAttachments() {
		attachments := execution.Attachments().List()
		for _, oneAttachment := range attachments {
			attachedValue := oneAttachment.Value()
			ins, err := app.executeValue(input, values, attachedValue)
			if err != nil {
				return nil, err
			}

			local := oneAttachment.Local()
			parameters[local] = ins
		}
	}

	execFn := module.Func()
	return execFn(parameters)
}
