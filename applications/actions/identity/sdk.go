package identity

import (
	"github.com/steve-care-software/syntax/applications/actions/identity/authenticates"
	"github.com/steve-care-software/syntax/applications/actions/identity/publics"
	"github.com/steve-care-software/syntax/domain/identity/wallets"
)

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithUnitRingSize(unitRingSize uint) Builder
	WithGenesisRingSize(genesisRingSize uint) Builder
	WithDepositTo(depositTo wallets.Wallet) Builder
	Now() (Application, error)
}

// Application represents an application
type Application interface {
	List() ([]string, error)
	New(name string, password []byte) error
	Public(name string) (publics.Application, error)
	Authenticate(name string, password []byte) (authenticates.Application, error)
}
