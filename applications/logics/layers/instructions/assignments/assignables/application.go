package assignables

import (
	"github.com/steve-care-software/datastencil/applications/logics/layers/instructions/assignments/assignables/bytes"
	"github.com/steve-care-software/datastencil/applications/logics/layers/instructions/assignments/assignables/compilers"
	"github.com/steve-care-software/datastencil/applications/logics/layers/instructions/assignments/assignables/constants"
	"github.com/steve-care-software/datastencil/applications/logics/layers/instructions/assignments/assignables/cryptography"
	"github.com/steve-care-software/datastencil/applications/logics/layers/instructions/assignments/assignables/databases"
	"github.com/steve-care-software/datastencil/applications/logics/layers/instructions/assignments/assignables/lists"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

type application struct {
	execCompilerApp compilers.Application
	execBytesApp    bytes.Application
	execConstantApp constants.Application
	execCryptoApp   cryptography.Application
	execDatabaseApp databases.Application
	execListApp     lists.Application
}

func createApplication(
	execCompilerApp compilers.Application,
	execBytesApp bytes.Application,
	execConstantApp constants.Application,
	execCryptoApp cryptography.Application,
	execDatabaseApp databases.Application,
	execListApp lists.Application,
) Application {
	out := application{
		execCompilerApp: execCompilerApp,
		execBytesApp:    execBytesApp,
		execConstantApp: execConstantApp,
		execCryptoApp:   execCryptoApp,
		execDatabaseApp: execDatabaseApp,
		execListApp:     execListApp,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, assignable assignables.Assignable) (stacks.Assignable, *uint, error) {
	if assignable.IsBytes() {
		bytesIns := assignable.Bytes()
		return app.execBytesApp.Execute(frame, bytesIns)
	}

	if assignable.IsConstant() {
		constant := assignable.Constant()
		return app.execConstantApp.Execute(constant)
	}

	if assignable.IsCryptography() {
		crypto := assignable.Cryptography()
		return app.execCryptoApp.Execute(frame, crypto)
	}

	if assignable.IsDatabase() {
		database := assignable.Database()
		return app.execDatabaseApp.Execute(frame, database)
	}

	if assignable.IsList() {
		list := assignable.List()
		return app.execListApp.Execute(frame, list)
	}

	compiler := assignable.Compiler()
	return app.execCompilerApp.Execute(frame, compiler)
}
