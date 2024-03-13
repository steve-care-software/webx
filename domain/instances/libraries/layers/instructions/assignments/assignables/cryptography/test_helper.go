package cryptography

import (
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/cryptography/decrypts"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/cryptography/encrypts"
)

// NewCryptographyWithDecryptForTests creates a new cryptography with decrypt for tests
func NewCryptographyWithDecryptForTests(decrypt decrypts.Decrypt) Cryptography {
	ins, err := NewBuilder().Create().WithDecrypt(decrypt).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewCryptographyWithEncryptForTests creates a new cryptography with encrypt for tests
func NewCryptographyWithEncryptForTests(encrypt encrypts.Encrypt) Cryptography {
	ins, err := NewBuilder().Create().WithEncrypt(encrypt).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
