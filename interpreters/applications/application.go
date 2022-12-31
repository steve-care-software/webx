package applications

import (
	"errors"
	"fmt"

	grammars_application "github.com/steve-care-software/webx/grammars/applications"
	"github.com/steve-care-software/webx/grammars/domain/grammars"
	programs_application "github.com/steve-care-software/webx/programs/applications"
	"github.com/steve-care-software/webx/programs/domain/instructions"
	"github.com/steve-care-software/webx/programs/domain/programs/modules"
	selectors_application "github.com/steve-care-software/webx/selectors/applications"
	"github.com/steve-care-software/webx/selectors/domain/selectors"
)

type application struct {
	grammarApp  grammars_application.Application
	selectorApp selectors_application.Application
	programApp  programs_application.Application
	grammarIns  grammars.Grammar
	selectorIns selectors.Selector
	modules     modules.Modules
}

func createApplication(
	grammarApp grammars_application.Application,
	selectorApp selectors_application.Application,
	programApp programs_application.Application,
	grammarIns grammars.Grammar,
	selectorIns selectors.Selector,
	modules modules.Modules,
) Application {
	out := application{
		grammarApp:  grammarApp,
		selectorApp: selectorApp,
		programApp:  programApp,
		grammarIns:  grammarIns,
		selectorIns: selectorIns,
		modules:     modules,
	}

	return &out
}

// ParseThenInterpret parses then interpret
func (app *application) ParseThenInterpret(input []interface{}, script []byte) ([]interface{}, []byte, error) {
	treeIns, err := app.grammarApp.Execute(app.grammarIns, script)
	if err != nil {
		str := fmt.Sprintf("there was an error while lexing the AST from the script using the grammar: %s", err.Error())
		return nil, nil, errors.New(str)
	}

	ins, isValid, remaining, err := app.selectorApp.Execute(app.selectorIns, treeIns)
	if err != nil {
		str := fmt.Sprintf("there was an error while parsing the instructions from the AST: %s", err.Error())
		return nil, nil, errors.New(str)
	}

	if !isValid {
		return nil, remaining, errors.New("the AST could not be parsed properly using the provided selector instance")
	}

	if castedInstructions, ok := ins.(instructions.Instructions); ok {
		programIns, err := app.programApp.Compile(app.modules, castedInstructions)
		if err != nil {
			str := fmt.Sprintf("there was an error while compiling the instructions into a program: %s", err.Error())
			return nil, remaining, errors.New(str)
		}

		output, err := app.programApp.Execute(input, programIns)
		if err != nil {
			str := fmt.Sprintf("there was an error while executing the compiled program: %s", err.Error())
			return nil, remaining, errors.New(str)
		}

		return output, remaining, nil
	}

	return nil, remaining, errors.New("the instructions returned by the selector's application could not be casted properly")
}
