package blockchains

import (
	"github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"
	"github.com/steve-care-software/webx/engine/units/domain/blockchains"
	"github.com/steve-care-software/webx/engine/units/domain/blockchains/blocks/transactions"
)

// Application represents the units database
type Application interface {
	List() ([]string, error)
	Insert(blockchain blockchains.Blockchain) error
	Update(currentIdentifier string, newIdentifier string) error
	Delete(hash hash.Hash) error
	Retrieve(identifier string) (blockchains.Blockchain, error)
	Transact(identifier string, trx transactions.Transactions) error
	Queue(identifier string) (transactions.Transactions, error)
	Mine(identifier string) error
	Sync(identifier string) error
}
