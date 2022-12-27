package compilers

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/databases/applications"
	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
	"github.com/steve-care-software/webx/roots/domain/compilers/compilers"
	"github.com/steve-care-software/webx/roots/domain/programs/programs"
	"github.com/steve-care-software/webx/roots/domain/programs/programs/modules"
	selector_applications "github.com/steve-care-software/webx/selectors/applications"
)

type application struct {
	databaseApp         applications.Application
	selectorSoftware    selector_applications.Application
	programBuilder      programs.Builder
	instructionsBuilder programs.InstructionsBuilder
	instructionBuilder  programs.InstructionBuilder
	applicationBuilder  programs.ApplicationBuilder
	attachmentsBuilder  programs.AttachmentsBuilder
	attachmentBuilder   programs.AttachmentBuilder
	valueBuilder        programs.ValueBuilder
}

func createApplication(
	databaseApp applications.Application,
	selectorSoftware selector_applications.Application,
	programBuilder programs.Builder,
	instructionsBuilder programs.InstructionsBuilder,
	instructionBuilder programs.InstructionBuilder,
	applicationBuilder programs.ApplicationBuilder,
	attachmentsBuilder programs.AttachmentsBuilder,
	attachmentBuilder programs.AttachmentBuilder,
	valueBuilder programs.ValueBuilder,
) Application {
	out := application{
		databaseApp:         databaseApp,
		selectorSoftware:    selectorSoftware,
		programBuilder:      programBuilder,
		instructionsBuilder: instructionsBuilder,
		instructionBuilder:  instructionBuilder,
		applicationBuilder:  applicationBuilder,
		attachmentsBuilder:  attachmentsBuilder,
		attachmentBuilder:   attachmentBuilder,
		valueBuilder:        valueBuilder,
	}

	return &out
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
	remaining := script
	instructionsList := []programs.Instruction{}
	executionsList := compiler.Executions().List()
	for index, oneExecution := range executionsList {
		appIns, rem, err := app.execute(uint(index), oneExecution, modules, remaining)
		if err != nil {
			return nil, err
		}

		instruction, err := app.instructionBuilder.Create().WithExecution(appIns).Now()
		if err != nil {
			return nil, err
		}

		instructionsList = append(instructionsList, instruction)
		if len(rem) <= 0 {
			break
		}

		remaining = rem
	}

	instructions, err := app.instructionsBuilder.Create().WithList(instructionsList).Now()
	if err != nil {
		return nil, err
	}

	builder := app.programBuilder.Create().WithInstructions(instructions)
	if compiler.HasOutputs() {
		outputs := compiler.Outputs()
		builder.WithOutputs(outputs)
	}

	return builder.Now()
}

func (app *application) execute(appIndex uint, execution compilers.Execution, modules modules.Modules, script []byte) (programs.Application, []byte, error) {
	remaining := script
	attachmentsList := []programs.Attachment{}
	parametersList := execution.Parameters().List()
	for idx, oneParameter := range parametersList {
		if len(remaining) <= 0 {
			break
		}

		index := oneParameter.Index()
		selector := oneParameter.Selector()
		valueIns, isValid, rem, err := app.selectorSoftware.Execute(selector, remaining)
		if err != nil {
			return nil, nil, err
		}

		if !isValid {
			str := fmt.Sprintf("the parameter (index: %d) could not be fetched properly by the selector application", idx)
			return nil, nil, errors.New(str)
		}

		if casted, ok := valueIns.([]byte); ok {
			value, err := app.valueBuilder.Create().WithConstant(casted).Now()
			if err != nil {
				return nil, nil, err
			}

			attachment, err := app.attachmentBuilder.Create().WithValue(value).WithLocal(index).Now()
			if err != nil {
				return nil, nil, err
			}

			attachmentsList = append(attachmentsList, attachment)
			remaining = rem
			continue
		}

		str := fmt.Sprintf("the selected value at parameter (index: %d) could notbe casted properly", idx)
		return nil, nil, errors.New(str)
	}

	attachments, err := app.attachmentsBuilder.Create().WithList(attachmentsList).Now()
	if err != nil {
		return nil, nil, err
	}

	moduleIndex := execution.ExecuteProgramModule()
	module, err := modules.Fetch(moduleIndex)
	if err != nil {
		return nil, nil, err
	}

	appIns, err := app.applicationBuilder.Create().WithIndex(appIndex).WithModule(module).WithAttachments(attachments).Now()
	if err != nil {
		return nil, nil, err
	}

	return appIns, remaining, nil
}
