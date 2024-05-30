package cryptography

import (
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/cryptography/decrypts"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/cryptography/encrypts"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/cryptography/keys"
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

// NewCryptographyWithKeyForTests creates a new cryptography with key for tests
func NewCryptographyWithKeyForTests(key keys.Key) Cryptography {
	ins, err := NewBuilder().Create().WithKey(key).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
