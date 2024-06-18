package cryptography

import (
	"github.com/steve-care-software/datastencil/applications/logics/instructions/assignments/assignables/cryptography/decrypts"
	"github.com/steve-care-software/datastencil/applications/logics/instructions/assignments/assignables/cryptography/encrypts"
	"github.com/steve-care-software/datastencil/applications/logics/instructions/assignments/assignables/cryptography/keys"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/cryptography"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// NewApplication creates a new application
func NewApplication(
	execDecryptApp decrypts.Application,
	execEncryptApp encrypts.Application,
	keyApp keys.Application,
) Application {
	return createApplication(
		execDecryptApp,
		execEncryptApp,
		keyApp,
	)
}

// Application represents an execution account application
type Application interface {
	Execute(frame stacks.Frame, assignable cryptography.Cryptography) (stacks.Assignable, *uint, error)
}
