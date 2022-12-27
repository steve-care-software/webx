package applications

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
	"github.com/steve-care-software/webx/programs/domain/programs"
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
	}
	return &out
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
