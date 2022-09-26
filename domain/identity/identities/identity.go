package identities

import (
	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/syntax/domain/identity/cryptography/signatures"
	"github.com/steve-care-software/syntax/domain/identity/identities/assets"
	"github.com/steve-care-software/syntax/domain/identity/identities/connections"
	"github.com/steve-care-software/syntax/domain/identity/identities/publics"
)

type identity struct {
	id          uuid.UUID
	public      publics.Public
	pk          signatures.PrivateKey
	connections connections.Connections
	assets      assets.Assets
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

func createIdentityWithAssets(
	id uuid.UUID,
	public publics.Public,
	pk signatures.PrivateKey,
	assets assets.Assets,
) Identity {
	return createIdentityInternally(id, public, pk, nil, assets)
}

func createIdentityWithConnectionsAndAssets(
	id uuid.UUID,
	public publics.Public,
	pk signatures.PrivateKey,
	connections connections.Connections,
	assets assets.Assets,
) Identity {
	return createIdentityInternally(id, public, pk, connections, assets)
}

func createIdentityInternally(
	id uuid.UUID,
	public publics.Public,
	pk signatures.PrivateKey,
	connections connections.Connections,
	assets assets.Assets,
) Identity {
	out := identity{
		id:          id,
		public:      public,
		pk:          pk,
		connections: connections,
		assets:      assets,
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

// HasAssets returns true if there is assets, false otherwise
func (obj *identity) HasAssets() bool {
	return obj.assets != nil
}

// Assets returns the assets, if any
func (obj *identity) Assets() assets.Assets {
	return obj.assets
}
