package authenticates

import (
	uuid "github.com/satori/go.uuid"
	identities "github.com/steve-care-software/syntax/domain/identity"
	"github.com/steve-care-software/syntax/domain/identity/cryptography/hash"
	"github.com/steve-care-software/syntax/domain/identity/publics"
	"github.com/steve-care-software/syntax/domain/identity/wallets"
)

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithPassword(password []byte) Builder
	WithUnitRingSize(unitRingSize uint) Builder
	WithGenesisRingSize(genesisRingSize uint) Builder
	Now() (Application, error)
}

// Application represents an identity application
type Application interface {
	Retrieve() (identities.Identity, error)
	Delete() error
	Connect(connection publics.Public) error
	Disconnect(id uuid.UUID) error
	Generate(depositTo wallets.Wallet, supply uint64, ticker string, description string) error
	Transfer(fromUnitHashes []hash.Hash, toOwner []hash.Hash, amount uint64, details string) error
}
