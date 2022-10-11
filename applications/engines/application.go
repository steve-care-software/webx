package engines

import (
	"github.com/steve-care-software/syntax/applications/engines/creates"
	"github.com/steve-care-software/syntax/applications/engines/criterias"
	"github.com/steve-care-software/syntax/applications/engines/grammars"
	"github.com/steve-care-software/syntax/applications/engines/interpreters"
	"github.com/steve-care-software/syntax/applications/engines/programs"
	program_application "github.com/steve-care-software/syntax/applications/engines/programs"
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

	progIns, remaining, err := app.programApp.Execute(grammar, command, script)
	if err != nil {
		return nil, nil, err
	}

	output, err := app.interpreterApp.Execute(input, progIns)
	if err != nil {
		return nil, nil, err
	}

	return output, remaining, nil
}
