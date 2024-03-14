package encryptions

import (
	"github.com/steve-care-software/datastencil/domain/instances/libraries/links/layers/instructions/assignments/assignables/accounts/encryptions/decrypts"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/links/layers/instructions/assignments/assignables/accounts/encryptions/encrypts"
)

// NewEncryptionWithEncryptForTests creates a new encryption with encrypt for tests
func NewEncryptionWithEncryptForTests(encrypt encrypts.Encrypt) Encryption {
	ins, err := NewBuilder().Create().WithEncrypt(encrypt).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewEncryptionWithDecryptForTests creates a new encryption with decrypt for tests
func NewEncryptionWithDecryptForTests(decrypt decrypts.Decrypt) Encryption {
	ins, err := NewBuilder().Create().WithDecrypt(decrypt).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
