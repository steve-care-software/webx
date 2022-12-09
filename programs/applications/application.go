package applications

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/blockchains/applications"
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
	contents_programs "github.com/steve-care-software/webx/programs/domain/contents/programs"
	contents_applications "github.com/steve-care-software/webx/programs/domain/contents/programs/applications"
	contents_instructions "github.com/steve-care-software/webx/programs/domain/contents/programs/instructions"
	contents_values "github.com/steve-care-software/webx/programs/domain/contents/programs/values"
	"github.com/steve-care-software/webx/programs/domain/programs"
	"github.com/steve-care-software/webx/programs/domain/programs/modules"
)

type application struct {
	blockchainApp             applications.Application
	contentAdapter            contents_programs.Adapter
	contentApplicationAdapter contents_applications.Adapter
	contentInstructionAdapter contents_instructions.Adapter
	contentValueAdapter       contents_values.Adapter
	builder                   programs.Builder
	instructionBuilder        programs.InstructionBuilder
	assignmentBuilder         programs.AssignmentBuilder
	applicationBuilder        programs.ApplicationBuilder
	attachmentsBuilder        programs.AttachmentsBuilder
	attachmentBuilder         programs.AttachmentBuilder
	valueBuilder              programs.ValueBuilder
	hashAdapter               hash.Adapter
}

func createApplication(
	blockchainApp applications.Application,
	contentAdapter contents_programs.Adapter,
	contentApplicationAdapter contents_applications.Adapter,
	contentInstructionAdapter contents_instructions.Adapter,
	contentValueAdapter contents_values.Adapter,
	hashAdapter hash.Adapter,
) Application {
	out := application{
		blockchainApp:             blockchainApp,
		contentAdapter:            contentAdapter,
		contentApplicationAdapter: contentApplicationAdapter,
		contentInstructionAdapter: contentInstructionAdapter,
		contentValueAdapter:       contentValueAdapter,
		hashAdapter:               hashAdapter,
	}
	return &out
}

// New creates a new application
func (app *application) New(name string) error {
	return app.blockchainApp.New(name)
}

// Retrieve retrieves a program by hash
func (app *application) Retrieve(context uint, hash hash.Hash, modules modules.Modules) (programs.Program, error) {
	content, err := app.blockchainApp.ReadByHash(context, hash)
	if err != nil {
		return nil, err
	}

	contentProgramIns, err := app.contentAdapter.ToProgram(content)
	if err != nil {
		return nil, err
	}

	insHashes := contentProgramIns.Instructions()
	instructions, err := app.retrieveInstructions(context, insHashes, modules)
	if err != nil {
		return nil, err
	}

	outputs := [][]byte{}
	outputIndexes := contentProgramIns.Outputs()
	for _, oneIndex := range outputIndexes {
		outputs = append(outputs, []byte(fmt.Sprintf("%d", oneIndex)))
	}

	return app.builder.Create().
		WithInstructions(instructions).
		WithOutputs(outputs).
		Now()
}

func (app *application) retrieveInstructions(context uint, hashes []hash.Hash, modules modules.Modules) ([]programs.Instruction, error) {
	output := []programs.Instruction{}
	for _, oneHash := range hashes {
		ins, err := app.retrieveInstruction(context, oneHash, modules)
		if err != nil {
			return nil, err
		}

		output = append(output, ins)
	}

	return output, nil
}

func (app *application) retrieveInstruction(context uint, hash hash.Hash, modules modules.Modules) (programs.Instruction, error) {
	content, err := app.blockchainApp.ReadByHash(context, hash)
	if err != nil {
		return nil, err
	}

	contentInstruction, err := app.contentInstructionAdapter.ToInstruction(content)
	if err != nil {
		return nil, err
	}

	contentIns := contentInstruction.Content()
	builder := app.instructionBuilder.Create()
	if contentIns.IsAssignment() {
		pAssignment := contentIns.Assignment()
		assignment, err := app.retrieveAssignment(context, *pAssignment, modules)
		if err != nil {
			return nil, err
		}

		builder.WithAssignment(assignment)
	}

	if contentIns.IsExecution() {
		pExecutionHash := contentIns.Execution()
		application, err := app.retrieveApplication(context, *pExecutionHash, modules)
		if err != nil {
			return nil, err
		}

		builder.WithExecution(application)
	}

	return builder.Now()
}

func (app *application) retrieveApplication(context uint, hash hash.Hash, modules modules.Modules) (programs.Application, error) {
	content, err := app.blockchainApp.ReadByHash(context, hash)
	if err != nil {
		return nil, err
	}

	contentApplication, err := app.contentApplicationAdapter.ToApplication(content)
	if err != nil {
		return nil, err
	}

	index := contentApplication.Module()
	module, err := modules.FetchByIndex(index)
	if err != nil {
		return nil, err
	}

	name := hash.Bytes()
	builder := app.applicationBuilder.Create().WithName(name).WithModule(module)
	if contentApplication.HasAttachments() {
		attachmentsList := []programs.Attachment{}
		contentAttachmentsList := contentApplication.Attachments().List()
		for _, oneContentAttachment := range contentAttachmentsList {
			valueHash := oneContentAttachment.Value()
			value, err := app.retrieveValue(context, valueHash, modules)
			if err != nil {
				return nil, err
			}

			localIndex := oneContentAttachment.Local()
			local := []byte(fmt.Sprintf("%d", localIndex))
			attachment, err := app.attachmentBuilder.Create().WithValue(value).WithLocal(local).Now()
			if err != nil {
				return nil, err
			}

			attachmentsList = append(attachmentsList, attachment)
		}

		attachments, err := app.attachmentsBuilder.Create().WithList(attachmentsList).Now()
		if err != nil {
			return nil, err
		}

		builder.WithAttachments(attachments)
	}

	return builder.Now()
}

func (app *application) retrieveAssignment(context uint, hash hash.Hash, modules modules.Modules) (programs.Assignment, error) {
	value, err := app.retrieveValue(context, hash, modules)
	if err != nil {
		return nil, err
	}

	name := hash.Bytes()
	return app.assignmentBuilder.Create().
		WithName(name).
		WithValue(value).
		Now()
}

func (app *application) retrieveValue(context uint, hash hash.Hash, modules modules.Modules) (programs.Value, error) {
	content, err := app.blockchainApp.ReadByHash(context, hash)
	if err != nil {
		return nil, err
	}

	contentValue, err := app.contentValueAdapter.ToValue(content)
	if err != nil {
		return nil, err
	}

	contentIns := contentValue.Content()
	builder := app.valueBuilder.Create()
	if contentIns.IsInput() {
		inputIndex := contentIns.Input()
		input := []byte(fmt.Sprintf("%d", inputIndex))
		builder.WithInput(input)
	}

	if contentIns.IsAssignment() {
		pAssignmentHash := contentIns.Assignment()
		assignment, err := app.retrieveAssignment(context, *pAssignmentHash, modules)
		if err != nil {
			return nil, err
		}

		builder.WithAssignment(assignment)
	}

	if contentIns.IsExecution() {
		pExecutionHash := contentIns.Execution()
		application, err := app.retrieveApplication(context, *pExecutionHash, modules)
		if err != nil {
			return nil, err
		}

		builder.WithExecution(application)
	}

	if contentIns.IsProgram() {
		pProgramHash := contentIns.Program()
		program, err := app.Retrieve(context, *pProgramHash, modules)
		if err != nil {
			return nil, err
		}

		builder.WithProgram(program)
	}

	if contentIns.IsConstant() {
		constant := contentIns.Constant()
		builder.WithConstant(constant)
	}

	return builder.Now()
}

// Scan scans the database for a program that can receive a given input and returns the requested output
func (app *application) Scan(context uint, input map[string]interface{}, callbackFn ScanCallbackFn) (programs.Program, error) {
	// retrieve the list of programs:

	// for each program, execute it with the input:

	// pass the output to the callback, if it returns true, keep it, then return the program with the smallest amount of points:
	return nil, nil
}

// Insert inserts a program
func (app *application) Insert(context uint, program programs.Program) error {
	return nil
}

// Execute executes a program
func (app *application) Execute(input map[string]interface{}, program programs.Program) (map[string]interface{}, error) {
	values := map[string]interface{}{}
	instructions := program.Instructions()
	for idx, oneInstruction := range instructions {
		if oneInstruction.IsAssignment() {
			assignment := oneInstruction.Assignment()
			name := assignment.Name()
			value := assignment.Value()

			ins, err := app.executeValue(input, values, value)
			if err != nil {
				str := fmt.Sprintf("there was an error while executing an assignment (index: %d. name: %s): %s", idx, name, err.Error())
				return nil, errors.New(str)
			}

			pHash, err := app.hashAdapter.FromBytes(name)
			if err != nil {
				return nil, err
			}

			values[pHash.String()] = ins
			continue
		}

		execution := oneInstruction.Execution()
		_, err := app.execute(input, values, execution)
		if err != nil {
			name := execution.Name()
			str := fmt.Sprintf("there was an error while executing an application (index: %d. name: %s): %s", idx, name, err.Error())
			return nil, errors.New(str)
		}
	}

	filtered := map[string]interface{}{}
	if program.HasOutputs() {
		outputs := program.Outputs()
		for _, oneOutput := range outputs {
			pHash, err := app.hashAdapter.FromBytes(oneOutput)
			if err != nil {
				return nil, err
			}

			keyname := pHash.String()
			if ins, ok := values[keyname]; ok {
				filtered[keyname] = ins
				continue
			}

			str := fmt.Sprintf("the program has an output parameter (name: %v, hash: %s), but the executed program does not contain that value", oneOutput, pHash.String())
			return nil, errors.New(str)
		}
	}

	return filtered, nil
}

func (app *application) executeValue(input map[string]interface{}, values map[string]interface{}, value programs.Value) (interface{}, error) {
	if value.IsInput() {
		inputName := value.Input()
		pHash, err := app.hashAdapter.FromBytes(inputName)
		if err != nil {
			return nil, err
		}

		keyname := pHash.String()
		if ins, ok := input[keyname]; ok {
			return ins, nil
		}

		str := fmt.Sprintf("the requested input variable (name: %s, hash: %s) is undefined", inputName, keyname)
		return nil, errors.New(str)
	}

	if value.IsAssignment() {
		assignment := value.Assignment()
		assignmentName := assignment.Name()
		pHash, err := app.hashAdapter.FromBytes(assignmentName)
		if err != nil {
			return nil, err
		}

		keyname := pHash.String()
		if ins, ok := values[keyname]; ok {
			return ins, nil
		}

		str := fmt.Sprintf("the requested variable (name: %v, hash: %s) is undefined", assignmentName, keyname)
		return nil, errors.New(str)
	}

	if value.IsConstant() {
		return value.Constant(), nil
	}

	if value.IsProgram() {
		subProgram := value.Program()
		subProgramOutput, err := app.Execute(input, subProgram)
		if err != nil {
			return nil, err
		}

		return subProgramOutput, nil
	}

	execution := value.Execution()
	return app.execute(input, values, execution)
}

func (app *application) execute(input map[string]interface{}, values map[string]interface{}, execution programs.Application) (interface{}, error) {
	module := execution.Module()
	parameters := map[string]interface{}{}
	if execution.HasAttachments() {
		attachments := execution.Attachments().List()
		for _, oneAttachment := range attachments {
			attachedValue := oneAttachment.Value()
			ins, err := app.executeValue(input, values, attachedValue)
			if err != nil {
				return nil, err
			}

			local := oneAttachment.Local()
			pHash, err := app.hashAdapter.FromBytes(local)
			if err != nil {
				return nil, err
			}

			parameters[pHash.String()] = ins
		}
	}

	execFn := module.Func()
	return execFn(parameters)
}
