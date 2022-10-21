package engines

import (
	"github.com/steve-care-software/webx/applications/creates"
	"github.com/steve-care-software/webx/applications/criterias"
	"github.com/steve-care-software/webx/applications/grammars"
	"github.com/steve-care-software/webx/applications/interpreters"
	"github.com/steve-care-software/webx/applications/programs"
	program_application "github.com/steve-care-software/webx/applications/programs"
	"github.com/steve-care-software/webx/domain/compilers"
)

type application struct {
	createApp      creates.Application
	criteriaApp    criterias.Application
	grammarApp     grammars.Application
	interpreterApp interpreters.Application
	programApp     program_application.Application
}

func createApplication(
	createApp creates.Application,
	criteriaApp criterias.Application,
	grammarApp grammars.Application,
	interpreterApp interpreters.Application,
	programApp program_application.Application,
) Application {
	out := application{
		createApp:      createApp,
		criteriaApp:    criteriaApp,
		grammarApp:     grammarApp,
		interpreterApp: interpreterApp,
		programApp:     programApp,
	}

	return &out
}

// Create returns the create application
func (app *application) Create() creates.Application {
	return app.createApp
}

// Criteria returns the criteria application
func (app *application) Criteria() criterias.Application {
	return app.criteriaApp
}

// Grammar returns the grammar application
func (app *application) Grammar() grammars.Application {
	return app.grammarApp
}

// Interpreter returns the grammar application
func (app *application) Interpreter() interpreters.Application {
	return app.interpreterApp
}

// Program returns the program application
func (app *application) Program() programs.Application {
	return app.programApp
}

// ParseThenInterpret parses the script then interpret its program
func (app *application) ParseThenInterpret(input map[string]interface{}, script []byte) (map[string]interface{}, []byte, error) {
	grammar, err := app.createApp.Grammar().Execute()
	if err != nil {
		return nil, nil, err
	}

	command, err := app.createApp.Command().Execute()
	if err != nil {
		return nil, nil, err
	}

	programOutput, err := app.programApp.Execute(grammar, command, script)
	if err != nil {
		return nil, nil, err
	}

	progIns := programOutput.Program()
	output, err := app.interpreterApp.Execute(input, progIns)
	if err != nil {
		return nil, nil, err
	}

	return output, programOutput.Remaining(), nil
}

// CompileThenParseThenInterpret compiles then parses then interpret
func (app *application) CompileThenParseThenInterpret(input map[string]interface{}, compiler compilers.Compiler, script []byte) (map[string]interface{}, []byte, []byte, error) {
	return nil, nil, nil, nil
}
