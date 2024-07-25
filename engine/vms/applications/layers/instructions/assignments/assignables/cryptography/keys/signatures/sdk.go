package signatures

import (
	"github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs"
	"github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures"
	"github.com/steve-care-software/webx/engine/vms/domain/keys/signers"
	"github.com/steve-care-software/webx/engine/vms/domain/stacks"
)

func NewApplication(
	voteApp votes.Application,
	signApp signs.Application,
) Application {
	assignableBuilder := stacks.NewAssignableBuilder()
	signerFactory := signers.NewFactory()
	return createApplication(
		voteApp,
		signApp,
		assignableBuilder,
		signerFactory,
	)
}

// Application represents an encryption application
type Application interface {
	Execute(frame stacks.Frame, assignable signatures.Signature) (stacks.Assignable, *uint, error)
}
