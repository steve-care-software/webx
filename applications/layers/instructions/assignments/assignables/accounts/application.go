package accounts

import (
	application_execution_communications "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/accounts/communications"
	application_execution_credentials "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/accounts/credentials"
	application_execution_encryptions "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/accounts/encryptions"
	application_execution_retrieves "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/accounts/retrieves"
	"github.com/steve-care-software/datastencil/domain/accounts"
	assignables_accounts "github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/accounts"
	"github.com/steve-care-software/datastencil/domain/stacks"
	stacks_accounts "github.com/steve-care-software/datastencil/domain/stacks/accounts"
)

type application struct {
	execCommApp        application_execution_communications.Application
	execCredentialsApp application_execution_credentials.Application
	execEncryptionApp  application_execution_encryptions.Application
	execRetrieveApp    application_execution_retrieves.Application
	repository         accounts.Repository
	service            accounts.Service
	assignableBuilder  stacks.AssignableBuilder
}

func createApplication(
	execCommApp application_execution_communications.Application,
	execCredentialsApp application_execution_credentials.Application,
	execEncryptionApp application_execution_encryptions.Application,
	execRetrieveApp application_execution_retrieves.Application,
	repository accounts.Repository,
	service accounts.Service,
	assignableBuilder stacks.AssignableBuilder,
	accountBuilder stacks_accounts.Builder,
) Application {
	out := application{
		execCommApp:        execCommApp,
		execCredentialsApp: execCredentialsApp,
		execEncryptionApp:  execEncryptionApp,
		repository:         repository,
		service:            service,
		assignableBuilder:  assignableBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, assignable assignables_accounts.Account) (stacks.Assignable, *uint, error) {
	builder := app.assignableBuilder.Create()
	if assignable.IsList() {
		variable := assignable.List()
		password, err := frame.FetchBytes(variable)
		if err != nil {
			return nil, nil, err
		}

		strList, err := app.repository.List(password)
		if err != nil {
			return nil, nil, err
		}

		builder = builder.WithStringList(strList)
	}

	if assignable.IsCredentials() {
		assCredentials := assignable.Credentials()
		return app.execCredentialsApp.Execute(frame, assCredentials)
	}

	if assignable.IsRetrieve() {
		retrieve := assignable.Retrieve()
		return app.execRetrieveApp.Execute(frame, retrieve)
	}

	if assignable.IsCommunication() {
		comm := assignable.Communication()
		return app.execCommApp.Execute(frame, comm)
	}

	if assignable.IsEncryption() {
		encryption := assignable.Encryption()
		return app.execEncryptionApp.Execute(frame, encryption)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, nil, nil
}
