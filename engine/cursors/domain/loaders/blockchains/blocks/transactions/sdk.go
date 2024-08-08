package transactions

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/blockchains/blocks/transactions/commands"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/singles/keys/signers"
)

// Transaction represents a transaction
type Transaction interface {
	Hash() hash.Hash
	Command() commands.Command
	Signature() signers.Signature
}
