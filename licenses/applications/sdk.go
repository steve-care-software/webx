package applications

import (
	"github.com/steve-care-software/webx/blockchains/applications"
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
	"github.com/steve-care-software/webx/licenses/domain/licenses"
	"github.com/steve-care-software/webx/licenses/domain/transfers"
)

// Application represents a program application
type Application interface {
	New(name string) error
	applications.Database
	Database
}

// Database represents the program database application
type Database interface {
	List() ([]string, error)
	RetrieveByOwner(context uint, pubKey hash.Hash) (licenses.Licenses, error)
	RetrieveByProduct(context uint, product hash.Hash) (licenses.Licenses, error)
	Register(license licenses.License) error
	Transfer(transfer transfers.Transfer) error
}
