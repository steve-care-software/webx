package encrypts

import (
	"github.com/steve-care-software/datastencil/domain/encryptors"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/cryptography/encrypts"
	"github.com/steve-care-software/datastencil/domain/stacks"
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
