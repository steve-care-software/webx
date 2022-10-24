package interpreters

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/domain/programs"
)

type application struct {
}

func createApplication() Application {
	out := application{}
	return &out
}

// Execute executes the interpreter
func (app *application) Execute(input map[string]interface{}, program programs.Program) (map[string]interface{}, error) {
	values := map[string]interface{}{}
	instructions := program.Instructions()
	for idx, oneInstruction := range instructions {
		if oneInstruction.IsAssignment() {
			assignment := oneInstruction.Assignment()
			name := assignment.Name()
			value := assignment.Value()

			ins, err := app.value(input, values, value)
			if err != nil {
				str := fmt.Sprintf("there was an error while executing an assignment (index: %d. name: %s): %s", idx, name, err.Error())
				return nil, errors.New(str)
			}

			values[name] = ins
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
			if ins, ok := values[oneOutput]; ok {
				filtered[oneOutput] = ins
				continue
			}

			str := fmt.Sprintf("the program has an output parameter (name: %s), but the executed program does not contain that value", oneOutput)
			return nil, errors.New(str)
		}
	}

	return filtered, nil
}

func (app *application) value(input map[string]interface{}, values map[string]interface{}, value programs.Value) (interface{}, error) {
	if value.IsInput() {
		inputName := value.Input()
		if ins, ok := input[inputName]; ok {
			return ins, nil
		}

		str := fmt.Sprintf("the requested input (name: %s) is undefined", inputName)
		return nil, errors.New(str)
	}

	if value.IsString() {
		str := value.String()
		return str, nil
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
			ins, err := app.value(input, values, attachedValue)
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
