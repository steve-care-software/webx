package blockchains

import (
	"github.com/steve-care-software/webx/engine/databases/entities/domain/hash"
	"github.com/steve-care-software/webx/engine/units/domain/blockchains"
	"github.com/steve-care-software/webx/engine/units/domain/blockchains/blocks/transactions"
)

// Application represents the units database
type Application interface {
	Insert(context uint, blockchain blockchains.Blockchain) error
	Delete(context uint, hash hash.Hash) error
	Retrieve(context uint, identifier string) (blockchains.Blockchain, error)
	Transact(context uint, identifier string, trx transactions.Transactions) error
}
