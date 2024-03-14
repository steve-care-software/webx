package signs

import (
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/accounts/communications/signs"
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
	Execute(frame stacks.Frame, assignable signs.Sign) (stacks.Assignable, *uint, error)
}
