package signs

import (
	"github.com/steve-care-software/webx/engine/vms/applications/instructions/assignments/assignables/cryptography/keys/signatures/signs/creates"
	"github.com/steve-care-software/webx/engine/vms/applications/instructions/assignments/assignables/cryptography/keys/signatures/signs/validates"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs"
	"github.com/steve-care-software/webx/engine/vms/domain/stacks"
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
