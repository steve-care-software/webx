package assignables

import (
	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/bytes"
	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/compilers"
	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/constants"
	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/cryptography"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

type application struct {
	execCompilerApp compilers.Application
	execBytesApp    bytes.Application
	execConstantApp constants.Application
	execCryptoApp   cryptography.Application
}

func createApplication(
	execCompilerApp compilers.Application,
	execBytesApp bytes.Application,
	execConstantApp constants.Application,
	execCryptoApp cryptography.Application,
) Application {
	out := application{
		execCompilerApp: execCompilerApp,
		execBytesApp:    execBytesApp,
		execConstantApp: execConstantApp,
		execCryptoApp:   execCryptoApp,
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

	compiler := assignable.Compiler()
	return app.execCompilerApp.Execute(frame, compiler)
}
