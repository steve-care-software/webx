package applications

import (
	"github.com/steve-care-software/webx/applications/creates"
	"github.com/steve-care-software/webx/applications/criterias"
	"github.com/steve-care-software/webx/applications/grammars"
	"github.com/steve-care-software/webx/applications/instructions"
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
	instructionApp instructions.Application
}

func createApplication(
	createApp creates.Application,
	criteriaApp criterias.Application,
	grammarApp grammars.Application,
	interpreterApp interpreters.Application,
	programApp program_application.Application,
	instructionApp instructions.Application,
) Application {
	out := application{
		createApp:      createApp,
		criteriaApp:    criteriaApp,
		grammarApp:     grammarApp,
		interpreterApp: interpreterApp,
		programApp:     programApp,
		instructionApp: instructionApp,
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

// Instruction returns the instruction application
func (app *application) Instruction() instructions.Application {
	return app.instructionApp
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

	instructionsOutput, err := app.instructionApp.Execute(grammar, command, script)
	if err != nil {
		return nil, nil, err
	}

	instructions := instructionsOutput.Instructions()
	program, err := app.programApp.Execute(instructions)
	if err != nil {
		return nil, nil, err
	}

	output, err := app.interpreterApp.Execute(input, program)
	if err != nil {
		return nil, nil, err
	}

	return output, instructionsOutput.Remaining(), nil
}

// CompileThenParseThenInterpret compiles then parses then interpret
func (app *application) CompileThenParseThenInterpret(input map[string]interface{}, compiler compilers.Compiler, script []byte) (map[string]interface{}, []byte, []byte, error) {
	return nil, nil, nil, nil
}
