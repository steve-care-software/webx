package decrypts

import (
	"github.com/steve-care-software/webx/engine/stencils/domain/encryptors"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/assignments/assignables/cryptography/decrypts"
	"github.com/steve-care-software/webx/engine/stencils/domain/stacks"
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
	Execute(frame stacks.Frame, assignable decrypts.Decrypt) (stacks.Assignable, *uint, error)
}
