package encrypts

import (
	"github.com/steve-care-software/datastencil/domain/hash"
)

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
