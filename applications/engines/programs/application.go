package programs

import (
	"errors"
	"fmt"
	"sort"

	application_criteria "github.com/steve-care-software/syntax/applications/engines/criterias"
	grammar_application "github.com/steve-care-software/syntax/applications/engines/grammars"
	"github.com/steve-care-software/syntax/domain/syntax/bytes/criterias"
	"github.com/steve-care-software/syntax/domain/syntax/bytes/grammars"
	"github.com/steve-care-software/syntax/domain/syntax/bytes/trees"
	"github.com/steve-care-software/syntax/domain/syntax/commands"
	"github.com/steve-care-software/syntax/domain/syntax/programs"
	"github.com/steve-care-software/syntax/domain/syntax/programs/applications"
	"github.com/steve-care-software/syntax/domain/syntax/programs/applications/modules"
)

type application struct {
	grammarApp         grammar_application.Application
	criteriaApp        application_criteria.Application
	builder            programs.Builder
	applicationBuilder applications.Builder
	attachmentsBuilder applications.AttachmentsBuilder
	attachmentBuilder  applications.AttachmentBuilder
	assignmentBuilder  applications.AssignmentBuilder
	valueBuilder       applications.ValueBuilder
	modules            modules.Modules
}

func createApplication(
	grammarApp grammar_application.Application,
	criteriaApp application_criteria.Application,
	builder programs.Builder,
	applicationBuilder applications.Builder,
	attachmentsBuilder applications.AttachmentsBuilder,
	attachmentBuilder applications.AttachmentBuilder,
	assignmentBuilder applications.AssignmentBuilder,
	valueBuilder applications.ValueBuilder,
) Application {
	return createApplicationInternally(
		grammarApp,
		criteriaApp,
		builder,
		applicationBuilder,
		attachmentsBuilder,
		attachmentBuilder,
		assignmentBuilder,
		valueBuilder,
		nil,
	)
}

func createApplicationWithModules(
	grammarApp grammar_application.Application,
	criteriaApp application_criteria.Application,
	builder programs.Builder,
	applicationBuilder applications.Builder,
	attachmentsBuilder applications.AttachmentsBuilder,
	attachmentBuilder applications.AttachmentBuilder,
	assignmentBuilder applications.AssignmentBuilder,
	valueBuilder applications.ValueBuilder,
	modules modules.Modules,
) Application {
	return createApplicationInternally(
		grammarApp,
		criteriaApp,
		builder,
		applicationBuilder,
		attachmentsBuilder,
		attachmentBuilder,
		assignmentBuilder,
		valueBuilder,
		modules,
	)
}

func createApplicationInternally(
	grammarApp grammar_application.Application,
	criteriaApp application_criteria.Application,
	builder programs.Builder,
	applicationBuilder applications.Builder,
	attachmentsBuilder applications.AttachmentsBuilder,
	attachmentBuilder applications.AttachmentBuilder,
	assignmentBuilder applications.AssignmentBuilder,
	valueBuilder applications.ValueBuilder,
	modules modules.Modules,
) Application {
	out := application{
		grammarApp:         grammarApp,
		criteriaApp:        criteriaApp,
		builder:            builder,
		applicationBuilder: applicationBuilder,
		attachmentsBuilder: attachmentsBuilder,
		attachmentBuilder:  attachmentBuilder,
		assignmentBuilder:  assignmentBuilder,
		valueBuilder:       valueBuilder,
		modules:            modules,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(grammar grammars.Grammar, command commands.Command, script []byte) (programs.Program, []byte, error) {
	assignments, outputParams, remaining, err := app.instructions(grammar, command, script)
	if err != nil {
		return nil, nil, err
	}

	builder := app.builder.Create().WithAssignments(assignments)
	if len(outputParams) > 0 {
		builder.WithOutputs(outputParams)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, remaining, nil
}

func (app *application) instructions(grammar grammars.Grammar, command commands.Command, script []byte) ([]applications.Assignment, []string, []byte, error) {
	index := uint(0)
	loadedParameters := map[string]bool{}
	loadedApplications := map[string]applications.Application{}
	loadedModules := map[string]modules.Module{}
	loadedAssignments := map[string]applications.Assignment{}
	remaining := script

	for {
		tree, err := app.grammarApp.Execute(grammar, remaining)
		if err != nil {
			break
		}

		retModules, retApplications, retParameters, retAssignments, err := app.executeCommand(
			index,
			tree,
			command,
			loadedModules,
			loadedApplications,
			loadedParameters,
			loadedAssignments,
		)

		if err != nil {
			return nil, nil, nil, err
		}

		if !tree.HasRemaining() {
			break
		}

		index++
		remaining = tree.Remaining()
		loadedModules = retModules
		loadedParameters = retParameters
		loadedApplications = retApplications
		loadedAssignments = retAssignments
		continue
	}

	// outputs:
	outputs := []string{}
	for name, isInput := range loadedParameters {
		if isInput {
			continue
		}

		outputs = append(outputs, name)
	}

	// sort the assignments:
	assignments := map[int]applications.Assignment{}
	for _, oneAssignment := range loadedAssignments {
		index := oneAssignment.Index()
		assignments[int(index)] = oneAssignment
	}

	keys := make([]int, 0, len(assignments))
	for k := range assignments {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	sortedAssignments := []applications.Assignment{}
	for _, k := range keys {
		sortedAssignments = append(sortedAssignments, assignments[k])
	}

	// returns:
	return sortedAssignments, outputs, remaining, nil
}

func (app *application) executeCommand(
	index uint,
	tree trees.Tree,
	command commands.Command,
	inModules map[string]modules.Module,
	inApplications map[string]applications.Application,
	inParameters map[string]bool,
	inVariables map[string]applications.Assignment,
) (map[string]modules.Module, map[string]applications.Application, map[string]bool, map[string]applications.Assignment, error) {
	// module declaration:
	moduleDeclaration := command.ModuleDeclaration()
	retModules, err := app.module(tree, moduleDeclaration, inModules)
	if retModules != nil {
		if err != nil {
			return nil, nil, nil, nil, err
		}

		return retModules, inApplications, inParameters, inVariables, nil
	}

	// application declaration:
	applicationDeclaration := command.ApplicationDeclaration()
	retApplications, err := app.applicationDeclaration(tree, applicationDeclaration, inModules, inApplications)
	if retApplications != nil {
		if err != nil {
			return nil, nil, nil, nil, err
		}

		return inModules, retApplications, inParameters, inVariables, nil
	}

	// parameter declaration:
	parameterDeclaration := command.ParameterDeclaration()
	retParameters, err := app.parameterDeclaration(tree, parameterDeclaration, inParameters)
	if retParameters != nil {
		if err != nil {
			return nil, nil, nil, nil, err
		}

		return inModules, inApplications, retParameters, inVariables, nil
	}

	// variable assignment:
	variableAssignment := command.VariableAssignment()
	retAssignments, err := app.variableAssignment(index, tree, variableAssignment, inParameters, inVariables, inApplications)
	if retAssignments != nil {
		if err != nil {
			return nil, nil, nil, nil, err
		}

		return inModules, inApplications, inParameters, retAssignments, nil
	}

	// attachment:
	attachment := command.Attachment()
	retAppAfterAttachments, err := app.attachment(tree, attachment, inParameters, inVariables, inApplications)
	if retAppAfterAttachments != nil {
		if err != nil {
			return nil, nil, nil, nil, err
		}

		return inModules, retAppAfterAttachments, inParameters, inVariables, nil
	}

	// execution:
	execution := command.Execution()
	retAssignmentsAfterExecution, err := app.execute(index, tree, execution, inParameters, inVariables, inApplications)
	if retAssignmentsAfterExecution != nil {
		if err != nil {
			return nil, nil, nil, nil, err
		}

		return inModules, inApplications, inParameters, retAssignmentsAfterExecution, nil
	}

	str := fmt.Sprintf("the command (%s) is invalid", tree.Bytes(true))
	return nil, nil, nil, nil, errors.New(str)
}

func (app *application) execute(
	index uint,
	tree trees.Tree,
	execution commands.Execution,
	inParameters map[string]bool,
	inVariables map[string]applications.Assignment,
	inApplications map[string]applications.Application,
) (map[string]applications.Assignment, error) {
	applicationCriteria := execution.Application()
	application, err := app.criteriaApp.Execute(applicationCriteria, tree)
	if err != nil {
		return nil, nil
	}

	applicationName := string(application)
	if appIns, ok := inApplications[applicationName]; ok {
		value, err := app.valueBuilder.Create().WithExecution(appIns).Now()
		if err != nil {
			return nil, err
		}

		assigneeCriteria := execution.Assignee()
		assignee, err := app.criteriaApp.Execute(assigneeCriteria, tree)
		if err != nil {
			return nil, err
		}

		assigneeName := string(assignee)
		assignment, err := app.assignmentBuilder.Create().WithIndex(index).WithName(assigneeName).WithValue(value).Now()
		if err != nil {
			return nil, err
		}

		inVariables[assigneeName] = assignment
		return inVariables, nil
	}

	str := fmt.Sprintf("the application (name: %s) is undeclared and therefore cannot be used in an execution", applicationName)
	return nil, errors.New(str)
}

func (app *application) attachment(
	tree trees.Tree,
	attachment commands.Attachment,
	inParameters map[string]bool,
	inVariables map[string]applications.Assignment,
	inApplications map[string]applications.Application,
) (map[string]applications.Application, error) {
	globalCriteria := attachment.Global()
	global, err := app.value(tree, globalCriteria, inParameters, inVariables, inApplications)
	if err != nil {
		return nil, nil
	}

	localCriteria := attachment.Local()
	local, err := app.criteriaApp.Execute(localCriteria, tree)
	if err != nil {
		return nil, nil
	}

	applicationCriteria := attachment.Application()
	application, err := app.criteriaApp.Execute(applicationCriteria, tree)
	if err != nil {
		return nil, nil
	}

	applicationName := string(application)
	if appIns, ok := inApplications[applicationName]; ok {
		localStr := string(local)
		attachment, err := app.attachmentBuilder.Create().WithValue(global).WithLocal(localStr).Now()
		if err != nil {
			return nil, err
		}

		attachmentsList := []applications.Attachment{}
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

func (app *application) variableAssignment(
	index uint,
	tree trees.Tree,
	variableAssignment commands.VariableAssignment,
	inParameters map[string]bool,
	inVariables map[string]applications.Assignment,
	inApplications map[string]applications.Application,
) (map[string]applications.Assignment, error) {
	valueCriteria := variableAssignment.Value()
	value, err := app.value(tree, valueCriteria, inParameters, inVariables, inApplications)
	if err != nil {
		return nil, err
	}

	if value == nil && err == nil {
		return nil, nil
	}

	assigneeCriteria := variableAssignment.Assignee()
	assignee, err := app.criteriaApp.Execute(assigneeCriteria, tree)
	if err != nil {
		return nil, nil
	}

	assigneeStr := string(assignee)
	assignment, err := app.assignmentBuilder.Create().WithIndex(index).WithName(assigneeStr).WithValue(value).Now()
	if err != nil {
		return nil, err
	}

	inVariables[assigneeStr] = assignment
	return inVariables, nil
}

func (app *application) value(
	tree trees.Tree,
	criteria criterias.Criteria,
	inParameters map[string]bool,
	inVariables map[string]applications.Assignment,
	inApplications map[string]applications.Application,
) (applications.Value, error) {
	value, err := app.criteriaApp.Execute(criteria, tree)
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

	return valueBuilder.Now()
}

func (app *application) parameterDeclaration(
	tree trees.Tree,
	parameterDeclaration commands.ParameterDeclaration,
	parameters map[string]bool,
) (map[string]bool, error) {
	inputCriteria := parameterDeclaration.Input()
	inputName, err := app.criteriaApp.Execute(inputCriteria, tree)
	if err == nil {
		nameStr := string(inputName)
		if _, ok := parameters[nameStr]; ok {
			str := fmt.Sprintf("the parameter (name: %s, isInput: %t) is already declared", nameStr, true)
			return nil, errors.New(str)
		}

		parameters[nameStr] = true
		return parameters, nil
	}

	outputCriteria := parameterDeclaration.Output()
	outputName, err := app.criteriaApp.Execute(outputCriteria, tree)
	if err != nil {
		return nil, nil
	}

	nameStr := string(outputName)
	if _, ok := parameters[nameStr]; ok {
		str := fmt.Sprintf("the parameter (name: %s, isInput: %t) is already declared", nameStr, false)
		return nil, errors.New(str)
	}

	parameters[nameStr] = false
	return parameters, nil
}

func (app *application) applicationDeclaration(
	tree trees.Tree,
	applicationDeclaration commands.ApplicationDeclaration,
	modules map[string]modules.Module,
	applications map[string]applications.Application,
) (map[string]applications.Application, error) {
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
	if _, ok := applications[nameStr]; ok {
		str := fmt.Sprintf("the application (name: %s) is already declared", nameStr)
		return nil, errors.New(str)
	}

	moduleNameStr := string(moduleName)
	if _, ok := modules[moduleNameStr]; !ok {
		str := fmt.Sprintf("the module (name: %s) is undefined but used in the application declaration (name: %s)", moduleNameStr, nameStr)
		return nil, errors.New(str)
	}

	application, err := app.applicationBuilder.Create().WithName(nameStr).WithModule(modules[moduleNameStr]).Now()
	if err != nil {
		return nil, err
	}

	applications[nameStr] = application
	return applications, nil
}

func (app *application) module(
	tree trees.Tree,
	moduleDeclaration commands.ModuleDeclaration,
	modules map[string]modules.Module,
) (map[string]modules.Module, error) {
	if app.modules == nil {
		return nil, nil
	}

	moduleDeclarationName := moduleDeclaration.Name()
	moduleName, err := app.criteriaApp.Execute(moduleDeclarationName, tree)
	if err != nil {
		return nil, nil
	}

	module, err := app.modules.Fetch(string(moduleName))
	if err != nil {
		return nil, err
	}

	name := module.Name()
	if _, ok := modules[name]; ok {
		str := fmt.Sprintf("the module (name: %s) is already loaded", name)
		return nil, errors.New(str)
	}

	modules[name] = module
	return modules, nil
}
