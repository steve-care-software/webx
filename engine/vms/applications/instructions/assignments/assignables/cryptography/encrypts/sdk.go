package encrypts

import (
	"github.com/steve-care-software/webx/engine/vms/domain/encryptors"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography/encrypts"
	"github.com/steve-care-software/webx/engine/vms/domain/stacks"
)

// NewApplication creates a new application
func NewApplication(
	encryptor encryptors.Encryptor,
) Application {
	assignableBuilder := stacks.NewAssignableBuilder()
	return createApplication(
		encryptor,
		assignableBuilder,
	)
}

// Application represents an execution account application
type Application interface {
	Execute(frame stacks.Frame, assignable encrypts.Encrypt) (stacks.Assignable, *uint, error)
}
