package libraries

import (
	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/libraries/compilers"
	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/libraries/databases"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/libraries"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

type application struct {
	execCompilerApp compilers.Application
	execDatabaseApp databases.Application
}

func createApplication(
	execCompilerApp compilers.Application,
	execDatabaseApp databases.Application,
) Application {
	out := application{
		execCompilerApp: execCompilerApp,
		execDatabaseApp: execDatabaseApp,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, assignable libraries.Library) (stacks.Assignable, *uint, error) {
	if assignable.IsCompiler() {
		compiler := assignable.Compiler()
		return app.execCompilerApp.Execute(frame, compiler)
	}

	database := assignable.Database()
	return app.execDatabaseApp.Execute(frame, database)
}
