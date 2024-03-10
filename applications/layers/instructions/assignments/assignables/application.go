package assignables

import (
	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/accounts"
	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/bytes"
	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/constants"
	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/cryptography"
	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/libraries"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

type application struct {
	execAccountApp    accounts.Application
	execBytesApp      bytes.Application
	execConstantApp   constants.Application
	execCryptoApp     cryptography.Application
	execLibraryApp    libraries.Application
	assignableBuilder stacks.AssignableBuilder
}

func createApplication(
	execAccountApp accounts.Application,
	execBytesApp bytes.Application,
	execConstantApp constants.Application,
	execCryptoApp cryptography.Application,
	execLibraryApp libraries.Application,
	assignableBuilder stacks.AssignableBuilder,
) Application {
	out := application{
		execAccountApp:    execAccountApp,
		execBytesApp:      execBytesApp,
		execConstantApp:   execConstantApp,
		execCryptoApp:     execCryptoApp,
		execLibraryApp:    execLibraryApp,
		assignableBuilder: assignableBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, assignable assignables.Assignable) (stacks.Assignable, error) {
	if assignable.IsBytes() {
		bytesIns := assignable.Bytes()
		return app.execBytesApp.Execute(frame, bytesIns)
	}

	if assignable.IsConstant() {
		constant := assignable.Constant()
		return app.execConstantApp.Execute(frame, constant)
	}

	if assignable.IsCryptography() {
		crypto := assignable.Cryptography()
		return app.execCryptoApp.Execute(frame, crypto)
	}

	if assignable.IsLibrary() {
		library := assignable.Library()
		return app.execLibraryApp.Execute(frame, library)
	}

	if assignable.IsAccount() {
		account := assignable.Account()
		return app.execAccountApp.Execute(frame, account)
	}

	query := assignable.Query()
	return app.assignableBuilder.Create().
		WithQuery(query).
		Now()
}
