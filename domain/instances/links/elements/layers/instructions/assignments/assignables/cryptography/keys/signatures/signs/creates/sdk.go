package creates

import (
	"github.com/steve-care-software/datastencil/domain/hash"
)

// Builder represents a builder
type Builder interface {
	Create() Builder
	WithMessage(message string) Builder
	WithPrivateKey(pk string) Builder
	Now() (Create, error)
}

// Create represents a create
type Create interface {
	Hash() hash.Hash
	Message() string
	PrivateKey() string
}
