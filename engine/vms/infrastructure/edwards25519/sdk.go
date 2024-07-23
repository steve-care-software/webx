package edwards25519

import (
	"github.com/steve-care-software/webx/engine/vms/domain/encryptors"
	"go.dedis.ch/kyber/v3/group/edwards25519"
)

var curve = edwards25519.NewBlakeSHA256Ed25519()

// NewEncryptor creates a new encryptor
func NewEncryptor() encryptors.Encryptor {
	return createEncryptor()
}
