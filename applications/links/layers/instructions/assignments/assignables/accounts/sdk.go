package accounts

import (
	application_execution_communications "github.com/steve-care-software/datastencil/applications/links/layers/instructions/assignments/assignables/accounts/communications"
	application_execution_credentials "github.com/steve-care-software/datastencil/applications/links/layers/instructions/assignments/assignables/accounts/credentials"
	application_execution_encryptions "github.com/steve-care-software/datastencil/applications/links/layers/instructions/assignments/assignables/accounts/encryptions"
	application_execution_retrieves "github.com/steve-care-software/datastencil/applications/links/layers/instructions/assignments/assignables/accounts/retrieves"
	"github.com/steve-care-software/datastencil/domain/accounts"
	assignables_accounts "github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/accounts"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// NewApplication creates a new application for tests
func NewApplication(
	execCommApp application_execution_communications.Application,
	execCredentialsApp application_execution_credentials.Application,
	execEncryptionApp application_execution_encryptions.Application,
	execRetrieveApp application_execution_retrieves.Application,
	repository accounts.Repository,
) Application {
	assignableBuilder := stacks.NewAssignableBuilder()
	return createApplication(
		execCommApp,
		execCredentialsApp,
		execEncryptionApp,
		execRetrieveApp,
		repository,
		assignableBuilder,
	)
}

// Application represents an execution account application
type Application interface {
	Execute(frame stacks.Frame, assignable assignables_accounts.Account) (stacks.Assignable, *uint, error)
}
