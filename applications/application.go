package applications

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/applications/creates"
	"github.com/steve-care-software/webx/applications/grammars"
	"github.com/steve-care-software/webx/applications/interpreters"
	"github.com/steve-care-software/webx/applications/programs"
	program_application "github.com/steve-care-software/webx/applications/programs"
	selector_application "github.com/steve-care-software/webx/applications/selectors"
	"github.com/steve-care-software/webx/domain/compilers"
	"github.com/steve-care-software/webx/domain/instructions"
)

type application struct {
	createApp      creates.Application
	grammarApp     grammars.Application
	interpreterApp interpreters.Application
	programApp     program_application.Application
	selectorApp    selector_application.Application
}

func createApplication(
	createApp creates.Application,
	grammarApp grammars.Application,
	interpreterApp interpreters.Application,
	programApp program_application.Application,
	selectorApp selector_application.Application,
) Application {
	out := application{
		createApp:      createApp,
		grammarApp:     grammarApp,
		interpreterApp: interpreterApp,
		programApp:     programApp,
		selectorApp:    selectorApp,
	}

	return &out
}

// Create returns the create application
func (app *application) Create() creates.Application {
	return app.createApp
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

// Selector returns the selector application
func (app *application) Selector() selector_application.Application {
	return app.selectorApp
}

// ParseThenInterpret parses the script then interpret its program
func (app *application) ParseThenInterpret(input map[string]interface{}, script []byte) (map[string]interface{}, []byte, error) {
	grammar, err := app.createApp.Grammar().Execute()
	if err != nil {
		return nil, nil, err
	}

	selector, err := app.createApp.Selector().Execute()
	if err != nil {
		return nil, nil, err
	}

	tree, err := app.grammarApp.Execute(grammar, script)
	if err != nil {
		return nil, nil, err
	}

	ins, isValid, err := app.selectorApp.Execute(selector, tree)
	if !isValid {
		str := fmt.Sprintf("the selector could not fetch the data on the grammar instance because it does NOT match")
		return nil, nil, errors.New(str)
	}

	if err != nil {
		return nil, nil, err
	}

	if casted, ok := ins.(instructions.Instructions); ok {
		program, err := app.programApp.Execute(casted)
		if err != nil {
			return nil, nil, err
		}

		output, err := app.interpreterApp.Execute(input, program)
		if err != nil {
			return nil, nil, err
		}

		return output, tree.Remaining(), nil
	}

	return nil, nil, errors.New("the selector was expected to fetch an Instructions instance, but its return value could not be casted properly")

}

// CompileThenParseThenInterpret compiles then parses then interpret
func (app *application) CompileThenParseThenInterpret(input map[string]interface{}, compiler compilers.Compiler, script []byte) (map[string]interface{}, []byte, []byte, error) {
	return nil, nil, nil, nil
}
