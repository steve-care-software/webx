package deletes

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/keys/signers"
)

// Builder represents a delete builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithVote(vote signers.Vote) Builder
	Now() (Delete, error)
}

// Delete represents a delete transaction
type Delete interface {
	Hash() hash.Hash
	Name() string
	Vote() signers.Vote
}
