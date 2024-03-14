package assignables

import (
	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/accounts"
	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/bytes"
	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/compilers"
	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/constants"
	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/cryptography"
	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/databases"
	"github.com/steve-care-software/datastencil/applications/layers/instructions/failures"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

type application struct {
	execCompilerApp   compilers.Application
	execDatabaseApp   databases.Application
	execAccountApp    accounts.Application
	execBytesApp      bytes.Application
	execConstantApp   constants.Application
	execCryptoApp     cryptography.Application
	assignableBuilder stacks.AssignableBuilder
}

func createApplication(
	execCompilerApp compilers.Application,
	execDatabaseApp databases.Application,
	execAccountApp accounts.Application,
	execBytesApp bytes.Application,
	execConstantApp constants.Application,
	execCryptoApp cryptography.Application,
	assignableBuilder stacks.AssignableBuilder,
) Application {
	out := application{
		execCompilerApp:   execCompilerApp,
		execDatabaseApp:   execDatabaseApp,
		execAccountApp:    execAccountApp,
		execBytesApp:      execBytesApp,
		execConstantApp:   execConstantApp,
		execCryptoApp:     execCryptoApp,
		assignableBuilder: assignableBuilder,
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

	if assignable.IsAccount() {
		account := assignable.Account()
		return app.execAccountApp.Execute(frame, account)
	}

	if assignable.IsCompiler() {
		compiler := assignable.Compiler()
		return app.execCompilerApp.Execute(frame, compiler)
	}

	if assignable.IsDatabase() {
		database := assignable.Database()
		return app.execDatabaseApp.Execute(frame, database)
	}

	queryVar := assignable.Query()
	query, err := frame.FetchQuery(queryVar)
	if err != nil {
		code := failures.CouldNotFetchQueryFromFrame
		return nil, &code, err
	}

	ins, err := app.assignableBuilder.Create().
		WithQuery(query).
		Now()

	if err != nil {
		return nil, nil, err
	}

	return ins, nil, nil
}
