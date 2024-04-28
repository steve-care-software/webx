package encryptions

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/accounts/encryptions/decrypts"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/accounts/encryptions/encrypts"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Adapter represents the encryption adapter
type Adapter interface {
	ToBytes(ins Encryption) ([]byte, error)
	ToInstance(bytes []byte) (Encryption, error)
}

// Builder represents an encryption builder
type Builder interface {
	Create() Builder
	WithEncrypt(encrypt encrypts.Encrypt) Builder
	WithDecrypt(decrypt decrypts.Decrypt) Builder
	Now() (Encryption, error)
}

// Encryption represents an encryption
type Encryption interface {
	Hash() hash.Hash
	IsEncrypt() bool
	Encrypt() encrypts.Encrypt
	IsDecrypt() bool
	Decrypt() decrypts.Decrypt
}
