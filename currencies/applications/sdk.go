package applications

import (
	"math/big"

	"github.com/steve-care-software/webx/blockchains/applications"
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
	"github.com/steve-care-software/webx/currencies/domain/genesis"
	"github.com/steve-care-software/webx/currencies/domain/spends"
)

// Application represents a program application
type Application interface {
	New(name string, genesis genesis.Genesis) error
	applications.Database
	Database
}

// Database represents the program database application
type Database interface {
	Genesis(context uint) (genesis.Genesis, error)
	Spend(context uint, spend spends.Spend) error
	Spends(context uint, spends spends.Spends) error
	Retrieve(context uint, hash hash.Hash) (spends.Spend, error)
	SearchByOwner(context uint, pubKey hash.Hash) (spends.Spends, error)
	SearchBySpender(context uint, pubKey hash.Hash) (spends.Spends, error)
	Balance(context uint, pubKey hash.Hash) (big.Int, error)
}
