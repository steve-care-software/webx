package encrypts

import "github.com/steve-care-software/datastencil/domain/hash"

// Builder represents an encrypt builder
type Builder interface {
	Create() Builder
	WithMessage(message string) Builder
	WithAccount(account string) Builder
	Now() (Encrypt, error)
}

// Encrypt represents an encrypt
type Encrypt interface {
	Hash() hash.Hash
	Message() string
	Account() string
}
