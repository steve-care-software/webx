package programs

import (
	"github.com/steve-care-software/webx/databases/applications"
	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
	"github.com/steve-care-software/webx/programs/domain/programs"
	"github.com/steve-care-software/webx/programs/domain/programs/modules"
	contents_programs "github.com/steve-care-software/webx/roots/domain/programs"
	contents_applications "github.com/steve-care-software/webx/roots/domain/programs/applications"
	contents_instructions "github.com/steve-care-software/webx/roots/domain/programs/instructions"
	contents_values "github.com/steve-care-software/webx/roots/domain/programs/values"
)

type application struct {
	databaseApp               applications.Application
	contentAdapter            contents_programs.Adapter
	contentBuilder            contents_programs.Builder
	contentApplicationAdapter contents_applications.Adapter
	contentApplicationBuilder contents_applications.Builder
	contentInstructionAdapter contents_instructions.Adapter
	contentInstructionBuilder contents_instructions.Builder
	contentValueAdapter       contents_values.Adapter
	contentValueBuilder       contents_values.Builder
	builder                   programs.Builder
	instructionsBuilder       programs.InstructionsBuilder
	instructionBuilder        programs.InstructionBuilder
	applicationBuilder        programs.ApplicationBuilder
	attachmentsBuilder        programs.AttachmentsBuilder
	attachmentBuilder         programs.AttachmentBuilder
	valueBuilder              programs.ValueBuilder
}

func createApplication(
	databaseApp applications.Application,
	contentAdapter contents_programs.Adapter,
	contentBuilder contents_programs.Builder,
	contentApplicationAdapter contents_applications.Adapter,
	contentApplicationBuilder contents_applications.Builder,
	contentInstructionAdapter contents_instructions.Adapter,
	contentInstructionBuilder contents_instructions.Builder,
	contentValueAdapter contents_values.Adapter,
	contentValueBuilder contents_values.Builder,
	builder programs.Builder,
	instructionsBuilder programs.InstructionsBuilder,
	instructionBuilder programs.InstructionBuilder,
	applicationBuilder programs.ApplicationBuilder,
	attachmentsBuilder programs.AttachmentsBuilder,
	attachmentBuilder programs.AttachmentBuilder,
	valueBuilder programs.ValueBuilder,
) Application {
	out := application{
		databaseApp:               databaseApp,
		contentAdapter:            contentAdapter,
		contentBuilder:            contentBuilder,
		contentApplicationAdapter: contentApplicationAdapter,
		contentApplicationBuilder: contentApplicationBuilder,
		contentInstructionAdapter: contentInstructionAdapter,
		contentInstructionBuilder: contentInstructionBuilder,
		contentValueAdapter:       contentValueAdapter,
		contentValueBuilder:       contentValueBuilder,
		builder:                   builder,
		instructionsBuilder:       instructionsBuilder,
		instructionBuilder:        instructionBuilder,
		applicationBuilder:        applicationBuilder,
		attachmentsBuilder:        attachmentsBuilder,
		attachmentBuilder:         attachmentBuilder,
		valueBuilder:              valueBuilder,
	}
	return &out
}

// Retrieve retrieves a program by hash
func (app *application) Retrieve(context uint, hash hash.Hash, modules modules.Modules) (programs.Program, error) {
	content, err := app.databaseApp.ReadByHash(context, hash)
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

	outputs := []uint{}
	outputIndexes := contentProgramIns.Outputs()
	for _, oneIndex := range outputIndexes {
		outputs = append(outputs, oneIndex)
	}

	return app.builder.Create().
		WithInstructions(instructions).
		WithOutputs(outputs).
		Now()
}

func (app *application) retrieveInstructions(context uint, hashes []hash.Hash, modules modules.Modules) (programs.Instructions, error) {
	list := []programs.Instruction{}
	for _, oneHash := range hashes {
		ins, err := app.retrieveInstruction(context, oneHash, modules)
		if err != nil {
			return nil, err
		}

		list = append(list, ins)
	}

	return app.instructionsBuilder.Create().
		WithList(list).
		Now()
}

func (app *application) retrieveInstruction(context uint, hash hash.Hash, modules modules.Modules) (programs.Instruction, error) {
	content, err := app.databaseApp.ReadByHash(context, hash)
	if err != nil {
		return nil, err
	}

	contentInstruction, err := app.contentInstructionAdapter.ToInstruction(content)
	if err != nil {
		return nil, err
	}

	contentIns := contentInstruction.Content()
	builder := app.instructionBuilder.Create()
	if contentIns.IsValue() {
		pValueHash := contentIns.Value()
		value, err := app.retrieveValue(context, *pValueHash, modules)
		if err != nil {
			return nil, err
		}

		builder.WithValue(value)
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
	content, err := app.databaseApp.ReadByHash(context, hash)
	if err != nil {
		return nil, err
	}

	contentApplication, err := app.contentApplicationAdapter.ToApplication(content)
	if err != nil {
		return nil, err
	}

	moduleIndex := contentApplication.Module()
	module, err := modules.Fetch(moduleIndex)
	if err != nil {
		return nil, err
	}

	index := contentApplication.Index()
	builder := app.applicationBuilder.Create().WithIndex(index).WithModule(module)
	if contentApplication.HasAttachments() {
		attachmentsList := []programs.Attachment{}
		contentAttachmentsList := contentApplication.Attachments().List()
		for _, oneContentAttachment := range contentAttachmentsList {
			valueHash := oneContentAttachment.Value()
			value, err := app.retrieveValue(context, valueHash, modules)
			if err != nil {
				return nil, err
			}

			local := oneContentAttachment.Local()
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

func (app *application) retrieveValue(context uint, hash hash.Hash, modules modules.Modules) (programs.Value, error) {
	content, err := app.databaseApp.ReadByHash(context, hash)
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
		pIndex := contentIns.Input()
		builder.WithInput(*pIndex)
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

func (app *application) insertApplication(context uint, application programs.Application) error {
	return nil
}

func (app *application) insertInstruction(context uint, instruction programs.Instruction) error {
	return nil
}

func (app *application) insertValue(context uint, value programs.Value) error {
	// if the value already exists:
	/*valueHash := value.Hash()
	_, err := app.Retrieve(context, valueHash)
	if err == nil {
		return nil
	}*/

	return nil
}
