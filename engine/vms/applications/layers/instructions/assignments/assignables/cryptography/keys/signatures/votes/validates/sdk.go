package validates

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/signers"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/validates"
	"github.com/steve-care-software/webx/engine/vms/domain/stacks"
)

// NewApplication creates a new application
func NewApplication() Application {
	assignableBuilder := stacks.NewAssignableBuilder()
	voteAdapter := signers.NewVoteAdapter()
	return createApplication(
		assignableBuilder,
		voteAdapter,
	)
}

// Application represents a validate vote application
type Application interface {
	Execute(frame stacks.Frame, assignable validates.Validate) (stacks.Assignable, *uint, error)
}
