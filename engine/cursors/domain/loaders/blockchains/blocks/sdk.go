package blocks

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/blockchains/blocks/transactions"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/storages/delimiters"
)

// Block represents a block
type Block interface {
	Hash() hash.Hash
	Transactions() delimiters.Delimiter
	HasCurrent() bool
	Current() transactions.Transaction
}
