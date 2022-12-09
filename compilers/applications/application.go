package applications

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/blockchains/applications"
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
	"github.com/steve-care-software/webx/compilers/domain/compilers"
	"github.com/steve-care-software/webx/programs/domain/instructions"
	instruction_applications "github.com/steve-care-software/webx/programs/domain/instructions/applications"
	"github.com/steve-care-software/webx/programs/domain/instructions/attachments"
	"github.com/steve-care-software/webx/programs/domain/instructions/parameters"
	"github.com/steve-care-software/webx/programs/domain/programs"
	"github.com/steve-care-software/webx/programs/domain/programs/modules"
)

type application struct {
	blockchainApp      applications.Application
	builder            programs.Builder
	instructionBuilder programs.InstructionBuilder
	applicationBuilder programs.ApplicationBuilder
	attachmentsBuilder programs.AttachmentsBuilder
	attachmentBuilder  programs.AttachmentBuilder
	assignmentBuilder  programs.AssignmentBuilder
	valueBuilder       programs.ValueBuilder
	hashAdapter        hash.Adapter
}

func createApplication(
	blockchainApp applications.Application,
	builder programs.Builder,
	instructionBuilder programs.InstructionBuilder,
	applicationBuilder programs.ApplicationBuilder,
	attachmentsBuilder programs.AttachmentsBuilder,
	attachmentBuilder programs.AttachmentBuilder,
	assignmentBuilder programs.AssignmentBuilder,
	valueBuilder programs.ValueBuilder,
	hashAdapter hash.Adapter,
) Application {
	out := application{
		blockchainApp:      blockchainApp,
		builder:            builder,
		instructionBuilder: instructionBuilder,
		applicationBuilder: applicationBuilder,
		attachmentsBuilder: attachmentsBuilder,
		attachmentBuilder:  attachmentBuilder,
		assignmentBuilder:  assignmentBuilder,
		valueBuilder:       valueBuilder,
		hashAdapter:        hashAdapter,
	}
	return &out
}

// New creates a new application
func (app *application) New(name string) error {
	return app.blockchainApp.New(name)
}

// List lists the compilers
func (app *application) List(ontext uint) ([]hash.Hash, error) {
	return nil, nil
}

// Retrieve retrieves a compiler by hash
func (app *application) Retrieve(context uint, hash hash.Hash) (compilers.Compiler, error) {
	return nil, nil
}

// Insert inserts a compiler
func (app *application) Insert(context uint, compiler compilers.Compiler) error {
	return nil
}

// InsertAll inserts a list of compilers
func (app *application) InsertAll(context uint, compilers []compilers.Compiler) error {
	return nil
}

// Execute executes a compiler
func (app *application) Execute(compiler compilers.Compiler, modules modules.Modules, script []byte) (programs.Program, error) {
	return nil, nil
}

func (app *application) instructionsToProgram(loaded modules.Modules, instructions instructions.Instructions) (programs.Program, error) {
	list := instructions.List()
	inModules := map[string]modules.Module{}
	inApplications := map[string]programs.Application{}
	inParameters := map[string]parameters.Parameter{}
	inAssignments := map[string]programs.Assignment{}
	inInstructions := []programs.Instruction{}
	for idx, oneInstruction := range list {
		outModules, outApplications, outParameters, outAssignments, outInstructions, err := app.instruction(
			loaded,
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

	outputs := [][]byte{}
	for _, parameter := range inParameters {
		if parameter.IsInput() {
			continue
		}

		outputs = append(outputs, parameter.Name())
	}

	builder := app.builder.Create().WithInstructions(inInstructions)
	if len(outputs) > 0 {
		builder.WithOutputs(outputs)
	}

	return builder.Now()
}

func (app *application) instruction(
	loaded modules.Modules,
	index uint,
	instruction instructions.Instruction,
	inModules map[string]modules.Module,
	inApplications map[string]programs.Application,
	inParameters map[string]parameters.Parameter,
	inAssignments map[string]programs.Assignment,
	inInstructions []programs.Instruction,
) (map[string]modules.Module, map[string]programs.Application, map[string]parameters.Parameter, map[string]programs.Assignment, []programs.Instruction, error) {
	if instruction.IsModule() {
		name := instruction.Module()
		outModules, err := app.module(name, inModules, loaded)
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
		outAssignments, newAssignment, err := app.assignment(loaded, index, assignment, inParameters, inAssignments, inApplications)
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
	executionHash, err := app.hashAdapter.FromBytes(execution)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	executionKeyname := executionHash.String()
	instructionBuilder := app.instructionBuilder.Create()
	if executedApp, ok := inApplications[executionKeyname]; ok {
		instructionBuilder.WithExecution(executedApp)
	} else {
		str := fmt.Sprintf("the application's execution (index: %d, name: %v, hash: %s) is invalid because the application is undefined", index, execution, executionKeyname)
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
	inParameters map[string]parameters.Parameter,
	inAssignments map[string]programs.Assignment,
	inApplications map[string]programs.Application,
) (map[string]programs.Application, error) {
	variable := attachment.Variable()
	currentValue, err := app.attachmentValue(variable, inParameters, inAssignments)
	if err != nil {
		return nil, err
	}

	applicationName := attachment.Application()
	pApplicationNameHash, err := app.hashAdapter.FromBytes(applicationName)
	if err != nil {
		return nil, err
	}

	applicationNameKeyname := pApplicationNameHash.String()
	if appIns, ok := inApplications[applicationNameKeyname]; ok {
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

		pNameHash, err := app.hashAdapter.FromBytes(name)
		if err != nil {
			return nil, err
		}

		nameKeyname := pNameHash.String()
		inApplications[nameKeyname] = updatedAppIns
		return inApplications, nil
	}

	str := fmt.Sprintf("the application (name: %s) is undeclared and therefore cannot be used in an attachment", applicationName)
	return nil, errors.New(str)
}

func (app *application) attachmentValue(
	variable attachments.Variable,
	inParameters map[string]parameters.Parameter,
	inAssignments map[string]programs.Assignment,
) (programs.Value, error) {
	current := variable.Current()
	pHash, err := app.hashAdapter.FromBytes(current)
	if err != nil {
		return nil, err
	}

	keyname := pHash.String()
	if currentIns, ok := inAssignments[keyname]; ok {
		return currentIns.Value(), nil
	}

	if parameter, ok := inParameters[keyname]; ok {
		if !parameter.IsInput() {
			str := fmt.Sprintf("the output variable (name: %v, hash: %s) cannot be used in attachment", current, keyname)
			return nil, errors.New(str)
		}

		return app.valueBuilder.Create().WithInput(current).Now()
	}

	str := fmt.Sprintf("the current variable (name: %v, hash: %s) is undeclared and therefore cannot be used in an attachment", current, keyname)
	return nil, errors.New(str)
}

func (app *application) assignment(
	loaded modules.Modules,
	index uint,
	assignment instructions.Assignment,
	inParameters map[string]parameters.Parameter,
	inAssignments map[string]programs.Assignment,
	inApplications map[string]programs.Application,
) (map[string]programs.Assignment, programs.Assignment, error) {
	valueIns, err := app.value(loaded, index, assignment, inParameters, inAssignments, inApplications)
	if err != nil {
		return nil, nil, err
	}

	variable := assignment.Variable()
	assignmentIns, err := app.assignmentBuilder.Create().WithIndex(index).WithName(variable).WithValue(valueIns).Now()
	if err != nil {
		return nil, nil, err
	}

	pHash, err := app.hashAdapter.FromBytes(variable)
	if err != nil {
		return nil, nil, err
	}

	keyname := pHash.String()
	inAssignments[keyname] = assignmentIns
	return inAssignments, assignmentIns, nil
}

func (app *application) value(
	loaded modules.Modules,
	index uint,
	assignment instructions.Assignment,
	inParameters map[string]parameters.Parameter,
	inAssignments map[string]programs.Assignment,
	inApplications map[string]programs.Application,
) (programs.Value, error) {
	variable := assignment.Variable()
	pHash, err := app.hashAdapter.FromBytes(variable)
	if err != nil {
		return nil, err
	}

	keyname := pHash.String()
	value := assignment.Value()
	builder := app.valueBuilder.Create()
	if value.IsVariable() {
		assignedVariable := value.Variable()
		pAssignedVariableHash, err := app.hashAdapter.FromBytes(assignedVariable)
		if err != nil {
			return nil, err
		}

		assignedVariableKeyname := pAssignedVariableHash.String()
		if parameter, ok := inParameters[assignedVariableKeyname]; ok {
			if !parameter.IsInput() {
				str := fmt.Sprintf("the assignment (index: %d, variable: %v, hash: %s) is using an output variable (nme: %v, hash: %s) as value", index, variable, keyname, assignedVariable, assignedVariableKeyname)
				return nil, errors.New(str)
			}

			builder.WithInput(variable)
		} else if assignment, ok := inAssignments[assignedVariableKeyname]; ok {
			builder.WithAssignment(assignment)
		} else {
			str := fmt.Sprintf("the assignment (index: %d, variable: %v, hash: %s) is using an undefined variable (name: %v, hash: %s) as value", index, variable, keyname, assignedVariable, assignedVariableKeyname)
			return nil, errors.New(str)
		}
	}

	if value.IsConstant() {
		constant := value.Constant()
		builder.WithConstant(constant)
	}

	if value.IsInstructions() {
		subInstructions := value.Instructions()
		subProgram, err := app.instructionsToProgram(loaded, subInstructions)
		if err != nil {
			return nil, err
		}

		builder.WithProgram(subProgram)
	}

	if value.IsExecution() {
		execution := value.Execution()
		pExecutionHash, err := app.hashAdapter.FromBytes(execution)
		if err != nil {
			return nil, err
		}

		executionKeyname := pExecutionHash.String()
		if executedApp, ok := inApplications[executionKeyname]; ok {
			builder.WithExecution(executedApp)
		} else {
			str := fmt.Sprintf("the assignment (index: %d, variable: %v, hash: %s) is using an undefined application execution (name: %v, hash: %s) as value", index, variable, keyname, execution, executionKeyname)
			return nil, errors.New(str)
		}
	}

	return builder.Now()
}

func (app *application) parameter(
	parameter parameters.Parameter,
	inParameters map[string]parameters.Parameter,
) (map[string]parameters.Parameter, error) {
	name := parameter.Name()
	pHash, err := app.hashAdapter.FromBytes(name)
	if err != nil {
		return nil, err
	}

	keyname := pHash.String()
	if _, ok := inParameters[keyname]; ok {
		str := fmt.Sprintf("the parameter (name: %s, hash: %s, isInput: %t) is already declared", name, keyname, parameter.IsInput())
		return nil, errors.New(str)
	}

	inParameters[keyname] = parameter
	return inParameters, nil
}

func (app *application) application(
	application instruction_applications.Application,
	inModules map[string]modules.Module,
	inApplications map[string]programs.Application,
) (map[string]programs.Application, error) {
	name := application.Name()
	pHash, err := app.hashAdapter.FromBytes(name)
	if err != nil {
		return nil, err
	}

	keyname := pHash.String()
	if _, ok := inApplications[keyname]; ok {
		str := fmt.Sprintf("the application (name: %v, hash: %s) is already declared", name, keyname)
		return nil, errors.New(str)
	}

	module := application.Module()
	pModuleHash, err := app.hashAdapter.FromBytes(module)
	if err != nil {
		return nil, err
	}

	moduleKeyname := pModuleHash.String()
	if _, ok := inModules[moduleKeyname]; !ok {
		fmt.Printf("\n%s, %s, %s, %v\n", moduleKeyname, name, module, inModules)
		str := fmt.Sprintf("the module (name: %v, keyname: %s) is undefined but used in the application declaration (name: %v, hash: %s)", module, moduleKeyname, name, keyname)
		return nil, errors.New(str)
	}

	ins, err := app.applicationBuilder.Create().WithName(name).WithModule(inModules[moduleKeyname]).Now()
	if err != nil {
		return nil, err
	}

	inApplications[keyname] = ins
	return inApplications, nil
}

func (app *application) module(
	name []byte,
	modules map[string]modules.Module,
	loaded modules.Modules,
) (map[string]modules.Module, error) {
	pHash, err := app.hashAdapter.FromBytes(name)
	if err != nil {
		return nil, err
	}

	keyname := pHash.String()
	if loaded == nil {
		str := fmt.Sprintf("the module (name: %v, hash: %s) is undefined because there is zero (0) module loaded", name, keyname)
		return nil, errors.New(str)
	}

	module, err := loaded.Fetch(name)
	if err != nil {
		return nil, err
	}

	if _, ok := modules[keyname]; ok {
		str := fmt.Sprintf("the module (name: %v, hash: %s) is already loaded", name, keyname)
		return nil, errors.New(str)
	}

	modules[keyname] = module
	return modules, nil
}
