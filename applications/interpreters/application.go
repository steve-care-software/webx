package interpreters

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/domain/programs"
	"github.com/steve-care-software/webx/domain/programs/applications"
)

type application struct {
}

func createApplication() Application {
	out := application{}
	return &out
}

func (app *application) value(input map[string]interface{}, values map[string]interface{}, value applications.Value) (interface{}, error) {
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

	execution := value.Execution()
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

// Execute executes the interpreter
func (app *application) Execute(input map[string]interface{}, program programs.Program) (map[string]interface{}, error) {
	values := map[string]interface{}{}
	assignments := program.Assignments()
	for idx, oneAssignment := range assignments {
		name := oneAssignment.Name()
		value := oneAssignment.Value()

		ins, err := app.value(input, values, value)
		if err != nil {
			str := fmt.Sprintf("there was an error while executing an assignment (index: %d. name: %s): %s", idx, name, err.Error())
			return nil, errors.New(str)
		}

		values[name] = ins
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
