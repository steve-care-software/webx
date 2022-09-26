package authenticates

import (
	"github.com/steve-care-software/syntax/domain/identity/cryptography/hash"
	"github.com/steve-care-software/syntax/domain/identity/identities"
	"github.com/steve-care-software/syntax/domain/identity/identities/publics"
	"github.com/steve-care-software/syntax/domain/identity/units"
)

// TransferEventFn represents a transfer event func
type TransferEventFn func(unit units.Unit, to []hash.Hash) error

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithPassword(password string) Builder
	WithBeforeTransferEventFn(before TransferEventFn) Builder
	WithAfterTransferEventFn(after TransferEventFn) Builder
	Now() (Application, error)
}

// Application represents an identity application
type Application interface {
	Retrieve() (identities.Identity, error)
	Connect(connection publics.Public) error
	Disconnect(hash hash.Hash) error
	Generate(supply uint64, ticker string, description string) error
	UnitByTicker(ticker string) (units.Unit, error)
	Tickers() ([]string, error)
	Transfer(unit units.Unit, to []hash.Hash) error
	Receive(unit units.Unit) error
}
