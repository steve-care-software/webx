package transactions

import (
	"github.com/steve-care-software/webx/engine/databases/entities/domain/hash"
	"github.com/steve-care-software/webx/engine/units/domain/blockchains/blocks/transactions/contents"
	"github.com/steve-care-software/webx/engine/units/domain/units/clears"
)

// Builder represents the transactions builder
type Builder interface {
	Create() Builder
	WithList(list []Transaction) Builder
	Now() (Transactions, error)
}

// Transactions represents transactions
type Transactions interface {
	Hash() hash.Hash
	List() []Transaction
}

// Transaction represents a transaction
type Transaction interface {
	Hash() hash.Hash
	Content() contents.Content
	Fee() clears.Clears
}
