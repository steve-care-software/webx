package decrypts

import (
	"github.com/steve-care-software/datastencil/domain/encryptors"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/links/layers/instructions/assignments/assignables/cryptography/decrypts"
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
	Execute(frame stacks.Frame, assignable decrypts.Decrypt) (stacks.Assignable, *uint, error)
}
