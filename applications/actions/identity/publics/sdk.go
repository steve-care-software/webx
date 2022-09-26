package publics

import (
	"github.com/steve-care-software/syntax/domain/identity/pendings"
	"github.com/steve-care-software/syntax/domain/identity/publics"
	"github.com/steve-care-software/syntax/domain/identity/units"
	"github.com/steve-care-software/syntax/domain/identity/wallets"
)

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithUnitRingSize(unitRingSize uint) Builder
	Now() (Application, error)
}

// Application represents a public application
type Application interface {
	Identity() (publics.Public, error)
	Connections() (publics.Publics, error)
	Prepare() (pendings.Pending, error)
	Receive(depositTo wallets.Wallet, pending pendings.Pending, unit units.Unit) error
}
