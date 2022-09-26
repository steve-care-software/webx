package identity

import (
	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/syntax/domain/identity/connections"
	"github.com/steve-care-software/syntax/domain/identity/cryptography/signatures"
	"github.com/steve-care-software/syntax/domain/identity/publics"
	"github.com/steve-care-software/syntax/domain/identity/wallets"
)

type identity struct {
	id          uuid.UUID
	public      publics.Public
	pk          signatures.PrivateKey
	connections connections.Connections
	wallets     wallets.Wallets
}

func createIdentity(
	id uuid.UUID,
	public publics.Public,
	pk signatures.PrivateKey,
) Identity {
	return createIdentityInternally(id, public, pk, nil, nil)
}

func createIdentityWithConnections(
	id uuid.UUID,
	public publics.Public,
	pk signatures.PrivateKey,
	connections connections.Connections,
) Identity {
	return createIdentityInternally(id, public, pk, connections, nil)
}

func createIdentityWithWallets(
	id uuid.UUID,
	public publics.Public,
	pk signatures.PrivateKey,
	wallets wallets.Wallets,
) Identity {
	return createIdentityInternally(id, public, pk, nil, wallets)
}

func createIdentityWithConnectionsAndWallets(
	id uuid.UUID,
	public publics.Public,
	pk signatures.PrivateKey,
	connections connections.Connections,
	wallets wallets.Wallets,
) Identity {
	return createIdentityInternally(id, public, pk, connections, wallets)
}

func createIdentityInternally(
	id uuid.UUID,
	public publics.Public,
	pk signatures.PrivateKey,
	connections connections.Connections,
	wallets wallets.Wallets,
) Identity {
	out := identity{
		id:          id,
		public:      public,
		pk:          pk,
		connections: connections,
		wallets:     wallets,
	}

	return &out
}

// ID returns the id
func (obj *identity) ID() uuid.UUID {
	return obj.id
}

// Public returns the public
func (obj *identity) Public() publics.Public {
	return obj.public
}

// PrivateKey returns the signature's pk
func (obj *identity) PrivateKey() signatures.PrivateKey {
	return obj.pk
}

// HasConnections returns true if there is connections, false otherwise
func (obj *identity) HasConnections() bool {
	return obj.connections != nil
}

// Connections returns the connections, if any
func (obj *identity) Connections() connections.Connections {
	return obj.connections
}

// HasWallets returns true if there is wallets, false otherwise
func (obj *identity) HasWallets() bool {
	return obj.wallets != nil
}

// Wallets returns the wallets, if any
func (obj *identity) Wallets() wallets.Wallets {
	return obj.wallets
}
