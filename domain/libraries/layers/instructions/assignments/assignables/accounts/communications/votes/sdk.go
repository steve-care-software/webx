package votes

import "github.com/steve-care-software/datastencil/domain/hash"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Builder represents a vote builder
type Builder interface {
	Create() Builder
	WithMessage(message string) Builder
	WithRing(ring string) Builder
	WithAccount(account string) Builder
	Now() (Vote, error)
}

// Vote represents a vote
type Vote interface {
	Hash() hash.Hash
	Message() string
	Ring() string
	Account() string
}
