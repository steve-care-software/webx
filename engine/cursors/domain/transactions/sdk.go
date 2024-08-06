package transactions

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	"github.com/steve-care-software/webx/engine/cursors/domain/signers"
	"github.com/steve-care-software/webx/engine/cursors/domain/transactions/commands"
)

// Transaction represents a transaction
type Transaction interface {
	Hash() hash.Hash
	Command() commands.Command
	Signature() signers.Signature
}
