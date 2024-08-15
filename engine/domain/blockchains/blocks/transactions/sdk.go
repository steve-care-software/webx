package transactions

import (
	"github.com/steve-care-software/webx/engine/domain/blockchains/blocks/transactions/contents"
	"github.com/steve-care-software/webx/engine/domain/hash"
)

// Transactions represents transactions
type Transactions interface {
	Hash() hash.Hash
	List() []Transaction
}

// Transaction represents a transaction
type Transaction interface {
	Hash() hash.Hash
	Content() contents.Content
	//Vote() signers.Vote
	Signature() string
}
