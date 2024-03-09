package signs

import "github.com/steve-care-software/datastencil/domain/hash"

// Builder represents a sign builder
type Builder interface {
	Create() Builder
	WithMessage(message string) Builder
	WithAccount(account string) Builder
	Now() (Sign, error)
}

// Sign represenst a sign
type Sign interface {
	Hash() hash.Hash
	Message() string
	Account() string
}
