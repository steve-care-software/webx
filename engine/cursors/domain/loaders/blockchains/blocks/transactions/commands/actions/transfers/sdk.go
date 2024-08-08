package transfers

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/blockchains/blocks/transactions/commands/actions/transfers/contents"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes"
)

// Transfer represents a transfer
type Transfer interface {
	Hash() hash.Hash
	Content() contents.Content
	Vote() votes.Vote
}
