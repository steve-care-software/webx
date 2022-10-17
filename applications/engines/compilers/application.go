package compilers

import (
	programs_application "github.com/steve-care-software/syntax/applications/engines/programs"
	"github.com/steve-care-software/syntax/domain/syntax/compilers"
	"github.com/steve-care-software/syntax/domain/syntax/programs"
)

type application struct {
	programApp programs_application.Application
}

func createApplication(
	programApp programs_application.Application,
) Application {
	out := application{
		programApp: programApp,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(compiler compilers.Compiler, script []byte) (programs.Program, []byte, []byte, error) {
	return nil, nil, nil, nil
}
