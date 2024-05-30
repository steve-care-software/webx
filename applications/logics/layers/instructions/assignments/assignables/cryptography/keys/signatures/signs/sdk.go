package signs

import (
	"github.com/steve-care-software/datastencil/applications/logics/layers//instructions/assignments/assignables/cryptography/keys/signatures/signs/creates"
	"github.com/steve-care-software/datastencil/applications/logics/layers//instructions/assignments/assignables/cryptography/keys/signatures/signs/validates"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// NewApplication creates a new application
func NewApplication(
	createApp creates.Application,
	validateApp validates.Application,
) Application {
	return createApplication(
		createApp,
		validateApp,
	)
}

// Application represents a sign application
type Application interface {
	Execute(frame stacks.Frame, assignable signs.Sign) (stacks.Assignable, *uint, error)
}
