package encryptions

import (
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/cryptography/keys/encryptions/decrypts"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/cryptography/keys/encryptions/encrypts"
)

// NewEncryptionWithGeneratePrivateKeyForTests creates a new encryption with generatePrivateKey for tests
func NewEncryptionWithGeneratePrivateKeyForTests() Encryption {
	ins, err := NewBuilder().Create().IsGeneratePrivateKey().Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewEncryptionWithFetchPublicKeyForTests creates a new encryption with fetchPublicKey for tests
func NewEncryptionWithFetchPublicKeyForTests(fetchPublicKey string) Encryption {
	ins, err := NewBuilder().Create().WithFetchPublicKey(fetchPublicKey).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewEncryptionWithEncryptForTests creates an encryption with encrypt for tests
func NewEncryptionWithEncryptForTests(encrypt encrypts.Encrypt) Encryption {
	ins, err := NewBuilder().Create().WithEncrypt(encrypt).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewEncryptionWithDecryptForTests creates a new decrypt for tests
func NewEncryptionWithDecryptForTests(decrypt decrypts.Decrypt) Encryption {
	ins, err := NewBuilder().Create().WithDecrypt(decrypt).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
