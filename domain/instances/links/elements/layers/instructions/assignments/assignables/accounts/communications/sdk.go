package communications

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/accounts/communications/signs"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/accounts/communications/votes"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Adapter represents the communication adapter
type Adapter interface {
	ToBytes(ins Communication) ([]byte, error)
	ToInstance(bytes []byte) (Communication, error)
}

// Builder represents a communication builder
type Builder interface {
	Create() Builder
	WithSign(sign signs.Sign) Builder
	WithVote(vote votes.Vote) Builder
	WithGenerateRing(genRing string) Builder
	Now() (Communication, error)
}

// Communication represents a communication
type Communication interface {
	Hash() hash.Hash
	IsSign() bool
	Sign() signs.Sign
	IsVote() bool
	Vote() votes.Vote
	IsGenerateRing() bool
	GenerateRing() string
}
