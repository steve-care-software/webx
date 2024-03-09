package retrieves

import (
	applications_accounts "github.com/steve-care-software/datastencil/applications/accounts"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/accounts/retrieves"
	"github.com/steve-care-software/datastencil/domain/stacks"
	"github.com/steve-care-software/datastencil/domain/stacks/accounts"
)

type application struct {
	accountApp        applications_accounts.Application
	accountBuilder    accounts.Builder
	assignableBuilder stacks.AssignableBuilder
}

func createApplication(
	accountApp applications_accounts.Application,
	accountBuilder accounts.Builder,
	assignableBuilder stacks.AssignableBuilder,
) Application {
	out := application{
		accountApp:        accountApp,
		accountBuilder:    accountBuilder,
		assignableBuilder: assignableBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, assignable retrieves.Retrieve) (stacks.Assignable, error) {
	passVar := assignable.Password()
	password, err := frame.FetchBytes(passVar)
	if err != nil {
		return nil, err
	}

	credVar := assignable.Credentials()
	credentials, err := frame.FetchCredentials(credVar)
	if err != nil {
		return nil, err
	}

	accountIns, err := app.accountApp.Retrieve(password, credentials)
	if err != nil {
		return nil, err
	}

	account, err := app.accountBuilder.Create().
		WithAccount(accountIns).
		Now()

	if err != nil {
		return nil, err
	}

	return app.assignableBuilder.Create().
		WithAccount(account).
		Now()
}
