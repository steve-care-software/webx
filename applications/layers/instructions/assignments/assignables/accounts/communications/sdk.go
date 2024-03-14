package communications

import (
	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/accounts/communications/signs"
	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/accounts/communications/votes"
	"github.com/steve-care-software/datastencil/domain/accounts/signers"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/accounts/communications"
	"github.com/steve-care-software/datastencil/domain/stacks"
	"github.com/steve-care-software/datastencil/domain/stacks/accounts"
)

// NewApplication creates a new application
func NewApplication(
	execSignApp signs.Application,
	execVoteApp votes.Application,
) Application {
	signerFactory := signers.NewFactory()
	accountBuilder := accounts.NewBuilder()
	assignableBuilder := stacks.NewAssignableBuilder()
	return createApplication(
		execSignApp,
		execVoteApp,
		signerFactory,
		accountBuilder,
		assignableBuilder,
	)
}

// Application represents an execution account application
type Application interface {
	Execute(frame stacks.Frame, assignable communications.Communication) (stacks.Assignable, *uint, error)
}
