package validates

import (
	"github.com/steve-care-software/datastencil/domain/hash"
)

// Builder creates  anew validate builder
type Builder interface {
	Create() Builder
	WithVote(vote string) Builder
	WithMessage(msg string) Builder
	WithHashedRing(hashedRing string) Builder
	Now() (Validate, error)
}

// Validate represents a validate
type Validate interface {
	Hash() hash.Hash
	Vote() string
	Message() string
	HashedRing() string
}
