package communications

import (
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/accounts/communications/signs"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/accounts/communications/votes"
)

// Builder represents a communication builder
type Builder interface {
	Create() Builder
	WithSign(sign signs.Sign) Builder
	WithVote(vote votes.Vote) Builder
	WithGenerateRing(genRing uint) Builder
	Now() (Communication, error)
}

// Communication represents a communication
type Communication interface {
	IsSign() bool
	Sign() signs.Sign
	IsVote() bool
	Vote() votes.Vote
	IsGenerateRing() bool
	GenerateRing() uint
}
