package credentials

import (
	account_credentials "github.com/steve-care-software/datastencil/domain/accounts/credentials"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/links/layers/instructions/assignments/assignables/accounts/credentials"
	"github.com/steve-care-software/datastencil/domain/stacks"
	"github.com/steve-care-software/datastencil/domain/stacks/accounts"
)

// NewApplication creates a new application
func NewApplication() Application {
	credentialsBuilder := account_credentials.NewBuilder()
	accountBuilder := accounts.NewBuilder()
	assignableBuilder := stacks.NewAssignableBuilder()
	return createApplication(
		credentialsBuilder,
		accountBuilder,
		assignableBuilder,
	)
}

// Application represents an execution account application
type Application interface {
	Execute(frame stacks.Frame, assignable credentials.Credentials) (stacks.Assignable, *uint, error)
}
