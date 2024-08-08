package commands

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/blockchains/blocks/transactions/commands/actions"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/blockchains/blocks/transactions/commands/locks"
)

// Command represents a command
type Command interface {
	Hash() hash.Hash
	Actions() actions.Actions
	Lock() locks.Lock
	Fees() uint64
}
