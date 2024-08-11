package edwards25519

import (
	"github.com/steve-care-software/webx/engine/cursors/applications/encryptions"
	"go.dedis.ch/kyber/v3/group/edwards25519"
)

var curve = edwards25519.NewBlakeSHA256Ed25519()

// NewEncryptionApplication creates a new encryption application
func NewEncryptionApplication() encryptions.Application {
	return createEncryptionApplication()
}
