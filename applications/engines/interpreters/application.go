package interpreters

import (
	"errors"

	"github.com/steve-care-software/syntax/applications/engines/creates"
	grammar_application "github.com/steve-care-software/syntax/applications/engines/grammars"
	program_application "github.com/steve-care-software/syntax/applications/engines/programs"
	"github.com/steve-care-software/syntax/domain/syntax/outputs"
	"github.com/steve-care-software/syntax/domain/syntax/programs"
	"github.com/steve-care-software/syntax/domain/syntax/programs/instructions/applications/modules"
)

type application struct {
	programApp program_application.Application
	grammarApp grammar_application.Application
	create     creates.Application
}

func createApplication(
	programApp program_application.Application,
	grammarApp grammar_application.Application,
	create creates.Application,
) Application {
	out := application{
		programApp: programApp,
		grammarApp: grammarApp,
		create:     create,
	}

	return &out
}

// Execute executes the interpreter
func (app *application) Execute(input map[string]interface{}, script []byte) (outputs.Output, []byte, error) {
	if app.create == nil {
		return nil, nil, errors.New("the create application is mandatory in order to build an interpreter Application instance")
	}

	grammar, err := app.create.Grammar().Execute()
	if err != nil {
		return nil, nil, err
	}

	tree, remaining, err := app.grammarApp.Execute(grammar, script)
	if err != nil {
		return nil, nil, err
	}

	command, err := app.create.Command().Execute()
	if err != nil {
		return nil, nil, err
	}

	progIns, err := app.programApp.Execute(tree, command)
	if err != nil {
		return nil, nil, err
	}

	modules, err := app.create.Modules().Execute()
	if err != nil {
		return nil, nil, err
	}

	output, err := app.execute(input, progIns, modules)
	if err != nil {
		return nil, nil, err
	}

	return output, remaining, nil
}

func (app *application) execute(input map[string]interface{}, program programs.Program, modules modules.Modules) (outputs.Output, error) {
	return nil, nil
}
