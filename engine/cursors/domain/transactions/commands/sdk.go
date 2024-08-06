package commands

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	"github.com/steve-care-software/webx/engine/cursors/domain/transactions/commands/contents"
)

// Command represents a command
type Command interface {
	Hash() hash.Hash
	Contents() contents.Contents
	Wallet() hash.Hash
	Fees() uint64
}
