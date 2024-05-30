package signatures

import (
	"github.com/steve-care-software/datastencil/applications/logics/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs"
	"github.com/steve-care-software/datastencil/applications/logics/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/cryptography/keys/signatures"
	"github.com/steve-care-software/datastencil/domain/keys/signers"
	"github.com/steve-care-software/datastencil/domain/stacks"
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
