package deletes

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/keys/signers"
)

// Delete represents a delete transaction
type Delete interface {
	Hash() hash.Hash
	Name() string
	Vote() signers.Vote
}
