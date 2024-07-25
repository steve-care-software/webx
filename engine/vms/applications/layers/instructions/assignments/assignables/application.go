package assignables

import (
	"github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/bytes"
	"github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/compilers"
	"github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/constants"
	"github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/cryptography"
	executables "github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/excutables"
	"github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/executions"
	"github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/lists"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables"
	"github.com/steve-care-software/webx/engine/vms/domain/stacks"
	"github.com/steve-care-software/webx/engine/vms/domain/stacks/failures"
)

type application struct {
	execCompilerApp  compilers.Application
	execExecutionApp executions.Application
	execBytesApp     bytes.Application
	execConstantApp  constants.Application
	execCryptoApp    cryptography.Application
	execListApp      lists.Application
	execExcutableApp executables.Application
}

func createApplication(
	execCompilerApp compilers.Application,
	execExecutionApp executions.Application,
	execBytesApp bytes.Application,
	execConstantApp constants.Application,
	execCryptoApp cryptography.Application,
	execListApp lists.Application,
	execExcutableApp executables.Application,
) Application {
	out := application{
		execCompilerApp:  execCompilerApp,
		execExecutionApp: execExecutionApp,
		execBytesApp:     execBytesApp,
		execConstantApp:  execConstantApp,
		execCryptoApp:    execCryptoApp,
		execListApp:      execListApp,
		execExcutableApp: execExcutableApp,
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

	if assignable.IsList() {
		list := assignable.List()
		return app.execListApp.Execute(frame, list)
	}

	if assignable.IsExecution() {
		execution := assignable.Execution()
		return app.execExecutionApp.Execute(frame, execution)
	}

	if assignable.IsExecutable() {
		executable := assignable.Executable()
		return app.execExcutableApp.Execute(frame, executable)
	}

	if assignable.IsVariable() {
		variable := assignable.Variable()
		assignable, err := frame.Fetch(variable)
		if err != nil {
			code := failures.CouldNotFetchFromFrame
			return nil, &code, err
		}

		return assignable, nil, nil
	}

	compiler := assignable.Compiler()
	return app.execCompilerApp.Execute(frame, compiler)
}
