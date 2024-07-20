package encryptions

import (
	"github.com/steve-care-software/datastencil/stencils/applications/layers/instructions/assignments/assignables/cryptography/keys/encryptions/decrypts"
	"github.com/steve-care-software/datastencil/stencils/applications/layers/instructions/assignments/assignables/cryptography/keys/encryptions/encrypts"
	"github.com/steve-care-software/datastencil/stencils/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/encryptions"
	"github.com/steve-care-software/datastencil/stencils/domain/keys/encryptors"
	"github.com/steve-care-software/datastencil/stencils/domain/stacks"
)

// NewApplication creates a new application
func NewApplication(
	decryptApp decrypts.Application,
	encryptApp encrypts.Application,
	bitRate int,
) Application {
	assignableBuilder := stacks.NewAssignableBuilder()
	pkBuilder := encryptors.NewBuilder()
	return createApplication(
		decryptApp,
		encryptApp,
		assignableBuilder,
		pkBuilder,
		bitRate,
	)
}

// Application represents an encryption application
type Application interface {
	Execute(frame stacks.Frame, assignable encryptions.Encryption) (stacks.Assignable, *uint, error)
}
