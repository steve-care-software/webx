package encryptions

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/accounts/encryptions/decrypts"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/accounts/encryptions/encrypts"
)

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
