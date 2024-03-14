package votes

import (
	"github.com/steve-care-software/datastencil/domain/instances/libraries/links/layers/instructions/assignments/assignables/accounts/communications/votes"
	"github.com/steve-care-software/datastencil/domain/stacks"
	stacks_accounts "github.com/steve-care-software/datastencil/domain/stacks/accounts"
)

// NewApplication creates a new application
func NewApplication() Application {
	accountBuilder := stacks_accounts.NewBuilder()
	assignableBuilder := stacks.NewAssignableBuilder()
	return createApplication(
		accountBuilder,
		assignableBuilder,
	)
}

// Application represents an execution account application
type Application interface {
	Execute(frame stacks.Frame, assignable votes.Vote) (stacks.Assignable, *uint, error)
}
