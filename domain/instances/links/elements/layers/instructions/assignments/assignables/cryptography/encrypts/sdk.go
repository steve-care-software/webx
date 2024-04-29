package encrypts

import "github.com/steve-care-software/datastencil/domain/hash"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the encrypt adapter
type Adapter interface {
	ToBytes(ins Encrypt) ([]byte, error)
	ToInstance(bytes []byte) (Encrypt, error)
}

// Builder represents an encrypt builder
type Builder interface {
	Create() Builder
	WithMessage(message string) Builder
	WithPassword(password string) Builder
	Now() (Encrypt, error)
}

// Encrypt represents an encrypt
type Encrypt interface {
	Hash() hash.Hash
	Message() string
	Password() string
}
