package interpreters

import (
	"github.com/steve-care-software/syntax/domain/syntax/outputs"
	"github.com/steve-care-software/syntax/domain/syntax/programs"
	"github.com/steve-care-software/syntax/domain/syntax/programs/instructions/applications/modules"
)

type application struct {
	modules modules.Modules
	program programs.Program
}

func createApplication(
	modules modules.Modules,
	program programs.Program,
) Application {
	out := application{
		modules: modules,
		program: program,
	}

	return &out
}

// Execute executes the interpreter
func (app *application) Execute(input map[string]interface{}) (outputs.Output, error) {
	return nil, nil
}
