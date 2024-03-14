package retrieves

import (
	"github.com/steve-care-software/datastencil/applications/layers/instructions/failures"
	"github.com/steve-care-software/datastencil/domain/accounts"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/accounts/retrieves"
	"github.com/steve-care-software/datastencil/domain/stacks"
	stacks_accounts "github.com/steve-care-software/datastencil/domain/stacks/accounts"
)

type application struct {
	repository        accounts.Repository
	accountBuilder    stacks_accounts.Builder
	assignableBuilder stacks.AssignableBuilder
}

func createApplication(
	repository accounts.Repository,
	accountBuilder stacks_accounts.Builder,
	assignableBuilder stacks.AssignableBuilder,
) Application {
	out := application{
		repository:        repository,
		accountBuilder:    accountBuilder,
		assignableBuilder: assignableBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, assignable retrieves.Retrieve) (stacks.Assignable, *uint, error) {
	passVar := assignable.Password()
	password, err := frame.FetchBytes(passVar)
	if err != nil {
		code := failures.CouldNotFetchPasswordFromFrame
		return nil, &code, err
	}

	credVar := assignable.Credentials()
	credentials, err := frame.FetchCredentials(credVar)
	if err != nil {
		code := failures.CouldNotFetchCredentialsFromFrame
		return nil, &code, err
	}

	accountIns, err := app.repository.Retrieve(password, credentials)
	if err != nil {
		code := failures.CouldNotRetrieveAccountFromDatabase
		return nil, &code, err
	}

	account, err := app.accountBuilder.Create().
		WithAccount(accountIns).
		Now()

	if err != nil {
		return nil, nil, err
	}

	ins, err := app.assignableBuilder.Create().
		WithAccount(account).
		Now()

	if err != nil {
		return nil, nil, err
	}

	return ins, nil, nil
}
