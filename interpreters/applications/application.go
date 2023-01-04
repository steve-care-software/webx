package applications

import (
	"errors"
	"fmt"

	grammars_application "github.com/steve-care-software/webx/grammars/applications"
	"github.com/steve-care-software/webx/grammars/domain/grammars"
	"github.com/steve-care-software/webx/interpreters/domain/results"
	programs_application "github.com/steve-care-software/webx/programs/applications"
	"github.com/steve-care-software/webx/programs/domain/instructions"
	selectors_application "github.com/steve-care-software/webx/selectors/applications"
	"github.com/steve-care-software/webx/selectors/domain/selectors"
)

type application struct {
	resultBuilder results.Builder
	grammarApp    grammars_application.Application
	selectorApp   selectors_application.Application
	programApp    programs_application.Application
	grammarIns    grammars.Grammar
	selectorIns   selectors.Selector
	modulesFn     FetchModulesFn
}

func createApplication(
	resultBuilder results.Builder,
	grammarApp grammars_application.Application,
	selectorApp selectors_application.Application,
	programApp programs_application.Application,
	grammarIns grammars.Grammar,
	selectorIns selectors.Selector,
	modulesFn FetchModulesFn,
) Application {
	out := application{
		resultBuilder: resultBuilder,
		grammarApp:    grammarApp,
		selectorApp:   selectorApp,
		programApp:    programApp,
		grammarIns:    grammarIns,
		selectorIns:   selectorIns,
		modulesFn:     modulesFn,
	}

	return &out
}

// ParseThenInterpret parses then interpret
func (app *application) ParseThenInterpret(input []interface{}, script []byte) (results.Result, error) {
	treeIns, err := app.grammarApp.Execute(app.grammarIns, script)
	if err != nil {
		str := fmt.Sprintf("there was an error while lexing the AST from the script using the grammar: %s", err.Error())
		return nil, errors.New(str)
	}

	ins, isValid, remaining, err := app.selectorApp.Execute(app.selectorIns, treeIns)
	if err != nil {
		str := fmt.Sprintf("there was an error while parsing the instructions from the AST: %s", err.Error())
		return nil, errors.New(str)
	}

	builder := app.resultBuilder.Create()
	if isValid {
		builder.IsValid()
	}

	if remaining != nil {
		builder.WithRemaining(remaining)
	}

	if castedInstructions, ok := ins.(instructions.Instructions); ok {
		modules, err := app.modulesFn()
		if err != nil {
			return nil, err
		}

		programIns, err := app.programApp.Compile(modules, castedInstructions)
		if err != nil {
			str := fmt.Sprintf("there was an error while compiling the instructions into a program: %s", err.Error())
			return nil, errors.New(str)
		}

		output, err := app.programApp.Execute(input, programIns)
		if err != nil {
			str := fmt.Sprintf("there was an error while executing the compiled program: %s", err.Error())
			return nil, errors.New(str)
		}

		if len(output) > 0 {
			builder.WithValues(output)
		}

		return builder.Now()
	}

	return nil, errors.New("the instructions returned by the selector's application could not be casted properly")
}
