package credentials

import (
	"github.com/steve-care-software/datastencil/applications/layers/instructions/failures"
	account_credentials "github.com/steve-care-software/datastencil/domain/accounts/credentials"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/accounts/credentials"
	"github.com/steve-care-software/datastencil/domain/stacks"
	"github.com/steve-care-software/datastencil/domain/stacks/accounts"
)

type application struct {
	credentialsBuilder account_credentials.Builder
	accountBuilder     accounts.Builder
	assignableBuilder  stacks.AssignableBuilder
}

func createApplication(
	credentialsBuilder account_credentials.Builder,
	accountBuilder accounts.Builder,
	assignableBuilder stacks.AssignableBuilder,
) Application {
	out := application{
		credentialsBuilder: credentialsBuilder,
		accountBuilder:     accountBuilder,
		assignableBuilder:  assignableBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, assignable credentials.Credentials) (stacks.Assignable, *uint, error) {
	userVar := assignable.Username()
	userBytes, err := frame.FetchBytes(userVar)
	if err != nil {
		code := failures.CouldNotFetchUsernameFromFrame
		return nil, &code, err
	}

	passVar := assignable.Password()
	password, err := frame.FetchBytes(passVar)
	if err != nil {
		code := failures.CouldNotFetchPasswordFromFrame
		return nil, &code, err
	}

	username := string(userBytes)
	credentials, err := app.credentialsBuilder.Create().
		WithUsername(username).
		WithPassword(password).
		Now()

	if err != nil {
		return nil, nil, err
	}

	account, err := app.accountBuilder.Create().
		WithCredentials(credentials).
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
