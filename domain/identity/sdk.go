package identity

import (
	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/syntax/domain/identity/connections"
	"github.com/steve-care-software/syntax/domain/identity/cryptography/signatures"
	"github.com/steve-care-software/syntax/domain/identity/publics"
	"github.com/steve-care-software/syntax/domain/identity/wallets"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents an identity builder
type Builder interface {
	Create() Builder
	WithID(id uuid.UUID) Builder
	WithPublic(public publics.Public) Builder
	WithPrivateKey(pk signatures.PrivateKey) Builder
	WithConnections(connections connections.Connections) Builder
	WithWallets(wallets wallets.Wallets) Builder
	Now() (Identity, error)
}

// Identity represents an identity
type Identity interface {
	ID() uuid.UUID
	Public() publics.Public
	PrivateKey() signatures.PrivateKey
	HasConnections() bool
	Connections() connections.Connections
	HasWallets() bool
	Wallets() wallets.Wallets
}

// Repository represents an identity repository
type Repository interface {
	Retrieve(name string, password []byte) (Identity, error)
}

// Service represents an identity service
type Service interface {
	Save(identity Identity, password []byte) error
	Delete(identity Identity, password []byte) error
}
