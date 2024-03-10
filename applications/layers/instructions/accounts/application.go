package accounts

import (
	"github.com/steve-care-software/datastencil/applications/layers/instructions/accounts/inserts"
	"github.com/steve-care-software/datastencil/applications/layers/instructions/accounts/updates"
	"github.com/steve-care-software/datastencil/applications/layers/instructions/failures"
	"github.com/steve-care-software/datastencil/domain/accounts"
	instructions_accounts "github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/accounts"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

type application struct {
	execInsertApp inserts.Application
	execUpdateApp updates.Application
	service       accounts.Service
}

func createApplication(
	execInsertApp inserts.Application,
	execUpdateApp updates.Application,
	service accounts.Service,
) Application {
	out := application{
		execInsertApp: execInsertApp,
		execUpdateApp: execUpdateApp,
		service:       service,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, instruction instructions_accounts.Account) (*uint, error) {
	if instruction.IsInsert() {
		insert := instruction.Insert()
		return app.execInsertApp.Execute(frame, insert)
	}

	if instruction.IsUpdate() {
		update := instruction.Update()
		return app.execUpdateApp.Execute(frame, update)
	}

	credentialsVar := instruction.Delete()
	credentials, err := frame.FetchCredentials(credentialsVar)
	if err != nil {
		code := failures.CouldNotFetchCredentialsFromFrame
		return &code, err
	}

	err = app.service.Delete(credentials)
	if err != nil {
		code := failures.CouldNotDeleteAccount
		return &code, err
	}

	return nil, nil
}
