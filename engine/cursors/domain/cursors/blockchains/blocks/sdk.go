package blocks

import "github.com/steve-care-software/webx/engine/cursors/domain/cursors/blockchains/blocks/transactions"

// Block represents a block
type Block interface {
	Index() uint64
	HasTransaction() bool
	Transaction() transactions.Transaction
}
