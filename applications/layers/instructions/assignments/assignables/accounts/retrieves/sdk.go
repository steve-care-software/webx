package retrieves

import (
	"github.com/steve-care-software/datastencil/domain/accounts"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/accounts/retrieves"
	"github.com/steve-care-software/datastencil/domain/stacks"
	stacks_accounts "github.com/steve-care-software/datastencil/domain/stacks/accounts"
)

// NewApplication creates a new application
func NewApplication(
	repository accounts.Repository,
) Application {
	accountBuilder := stacks_accounts.NewBuilder()
	assignableBuilder := stacks.NewAssignableBuilder()
	return createApplication(
		repository,
		accountBuilder,
		assignableBuilder,
	)
}

// Application represents an execution account application
type Application interface {
	Execute(frame stacks.Frame, assignable retrieves.Retrieve) (stacks.Assignable, *uint, error)
}
