package creates

import (
	"github.com/steve-care-software/datastencil/domain/hash"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder creates a create builder
type Builder interface {
	Create() Builder
	WithMessage(message string) Builder
	WithRing(ring string) Builder
	WithPrivateKey(pk string) Builder
	Now() (Create, error)
}

// Create represents a cretae vote
type Create interface {
	Hash() hash.Hash
	Message() string
	Ring() string
	PrivateKey() string
}
