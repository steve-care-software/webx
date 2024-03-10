package cryptography

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/cryptography/decrypts"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/cryptography/encrypts"
)

// Builder represents a cryptography builder
type Builder interface {
	Create() Builder
	WithEncrypt(encrypt encrypts.Encrypt) Builder
	WithDecrypt(decrypt decrypts.Decrypt) Builder
	Now() (Cryptography, error)
}

// Cryptography represents a cryptography
type Cryptography interface {
	Hash() hash.Hash
	IsEncrypt() bool
	Encrypt() encrypts.Encrypt
	IsDecrypt() bool
	Decrypt() decrypts.Decrypt
}
