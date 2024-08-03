package blocks

import (
	"github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"
	"github.com/steve-care-software/webx/engine/units/domain/blockchains/blocks/transactions"
)

// Block represents a block
type Block interface {
	Hash() hash.Hash
	Transactions() transactions.Transactions
	Parent() hash.Hash
}
