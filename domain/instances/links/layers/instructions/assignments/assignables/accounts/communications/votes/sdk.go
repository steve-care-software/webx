package votes

import "github.com/steve-care-software/datastencil/domain/hash"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Adapter represents the vote adapter
type Adapter interface {
	ToBytes(ins Vote) ([]byte, error)
	ToInstance(bytes []byte) (Vote, error)
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
