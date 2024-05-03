package keys

import (
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/cryptography/keys/encryptions"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/cryptography/keys/signatures"
)

// NewKeyWithEncryptionForTests creates a new key with encryption for tests
func NewKeyWithEncryptionForTests(encryption encryptions.Encryption) Key {
	ins, err := NewBuilder().Create().WithEncryption(encryption).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewKeyWithSignatureForTests creates a new key with signature for tests
func NewKeyWithSignatureForTests(signature signatures.Signature) Key {
	ins, err := NewBuilder().Create().WithSignature(signature).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
