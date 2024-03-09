package assignables

import (
	"github.com/steve-care-software/datastencil/applications/libraries/executions/instructions/assignables/accounts"
	"github.com/steve-care-software/datastencil/applications/libraries/executions/instructions/assignables/bytes"
	"github.com/steve-care-software/datastencil/applications/libraries/executions/instructions/assignables/constants"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

type application struct {
	execAccountApp  accounts.Application
	execBytesApp    bytes.Application
	execConstantApp constants.Application
}

func createApplication(
	execAccountApp accounts.Application,
	execBytesApp bytes.Application,
	execConstantApp constants.Application,
) Application {
	out := application{
		execAccountApp:  execAccountApp,
		execBytesApp:    execBytesApp,
		execConstantApp: execConstantApp,
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

	account := assignable.Account()
	return app.execAccountApp.Execute(frame, account)
}
