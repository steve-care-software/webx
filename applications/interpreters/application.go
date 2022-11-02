package interpreters

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/domain/cryptography/hash"
	"github.com/steve-care-software/webx/domain/programs"
)

type application struct {
	hashAdapter hash.Adapter
}

func createApplication(
	hashAdapter hash.Adapter,
) Application {
	out := application{
		hashAdapter: hashAdapter,
	}

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

func (app *application) value(input map[string]interface{}, values map[string]interface{}, value programs.Value) (interface{}, error) {
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
			ins, err := app.value(input, values, attachedValue)
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
