package decrypts

import (
	"github.com/steve-care-software/historydb/domain/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the decrypt adapter
type Adapter interface {
	ToBytes(ins Decrypt) ([]byte, error)
	ToInstance(bytes []byte) (Decrypt, error)
}

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
