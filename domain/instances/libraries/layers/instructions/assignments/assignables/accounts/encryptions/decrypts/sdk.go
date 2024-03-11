package decrypts

import "github.com/steve-care-software/datastencil/domain/hash"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Builder represents a decrypt builder
type Builder interface {
	Create() Builder
	WithCipher(cipher string) Builder
	WithAccount(account string) Builder
	Now() (Decrypt, error)
}

// Decrypt represents a decrypt
type Decrypt interface {
	Hash() hash.Hash
	Cipher() string
	Account() string
}
