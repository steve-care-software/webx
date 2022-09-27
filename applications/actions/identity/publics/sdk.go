package publics

import (
	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/syntax/domain/identity/pendings"
	"github.com/steve-care-software/syntax/domain/identity/publics"
	"github.com/steve-care-software/syntax/domain/identity/publics/assets"
	"github.com/steve-care-software/syntax/domain/identity/publics/assets/units"
	"github.com/steve-care-software/syntax/domain/identity/wallets"
)

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithUnitRingSize(unitRingSize uint) Builder
	WithDepositTo(depositTo wallets.Wallet) Builder
	Now() (Application, error)
}

// Application represents a public application
type Application interface {
	Retrieve() (publics.Public, error)
	Prepare() (pendings.Pending, error)
	Receive(pending pendings.Pending, unit units.Unit) error
	Connect(connection publics.Public) error
	Assets(publicID uuid.UUID, newAssets assets.Assets) error
	Asset(publicID uuid.UUID, newAsset assets.Asset) error
}
