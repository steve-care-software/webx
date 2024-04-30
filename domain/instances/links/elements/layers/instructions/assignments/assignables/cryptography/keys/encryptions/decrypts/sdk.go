package decrypts

import (
	"github.com/steve-care-software/datastencil/domain/hash"
)

// Builder represents a decrypt builder
type Builder interface {
	Create() Builder
	WithCipher(cipher string) Builder
	WithPrivateKey(pk string) Builder
	Now() (Decrypt, error)
}

// Decrypt represents a decrypt
type Decrypt interface {
	Hash() hash.Hash
	Cipher() string
	PrivateKey() string
}
