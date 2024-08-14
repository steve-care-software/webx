package blocks

import (
	"github.com/steve-care-software/webx/engine/domain/blockchains/blocks/transactions"
	"github.com/steve-care-software/webx/engine/domain/blockchains/hash"
)

// Block represents a block
type Block interface {
	Hash() hash.Hash
	Transactions() transactions.Transactions
	Answer() []byte
	Difficulty() uint8 // the difficulty of the answer on the hash, do not store that value, calculate it
}
