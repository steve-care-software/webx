package encrypts

import (
	"github.com/steve-care-software/datastencil/states/domain/hash"
)

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
	WithMessage(msg string) Builder
	WithPublicKey(pubKey string) Builder
	Now() (Encrypt, error)
}

// Encrypt represents an encrypt
type Encrypt interface {
	Hash() hash.Hash
	Message() string
	PublicKey() string
}
