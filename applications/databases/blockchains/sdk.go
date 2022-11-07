package blockchains

import (
	"github.com/steve-care-software/webx/domain/blockchains/blocks"
	"github.com/steve-care-software/webx/domain/blockchains/transactions"
	"github.com/steve-care-software/webx/domain/cryptography/hash"
	"github.com/steve-care-software/webx/domain/databases"
)

// Builder represents the application builder
type Builder interface {
	Create() Builder
	WithDtabase(database databases.Database) Builder
	Now() (Application, error)
}

// Application represents the blockchain application
type Application interface {
	Reference() hash.Hash
	Pendings() (transactions.Transactions, error)
	Search(trx transactions.Transaction) (blocks.Block, error)
	Transact(trx transactions.Transaction) error
}
