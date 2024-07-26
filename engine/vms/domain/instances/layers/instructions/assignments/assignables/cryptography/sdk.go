package cryptography

import (
	"github.com/steve-care-software/webx/engine/databases/entities/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography/decrypts"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography/encrypts"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography/keys"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the cryptography adapter
type Adapter interface {
	ToBytes(ins Cryptography) ([]byte, error)
	ToInstance(bytes []byte) (Cryptography, error)
}

// Builder represents a cryptography builder
type Builder interface {
	Create() Builder
	WithEncrypt(encrypt encrypts.Encrypt) Builder
	WithDecrypt(decrypt decrypts.Decrypt) Builder
	WithKey(key keys.Key) Builder
	Now() (Cryptography, error)
}

// Cryptography represents a cryptography
type Cryptography interface {
	Hash() hash.Hash
	IsEncrypt() bool
	Encrypt() encrypts.Encrypt
	IsDecrypt() bool
	Decrypt() decrypts.Decrypt
	IsKey() bool
	Key() keys.Key
}
