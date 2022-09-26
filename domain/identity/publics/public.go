package publics

import (
	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/syntax/domain/identity/connections"
	"github.com/steve-care-software/syntax/domain/identity/cryptography/encryptions/keys"
	"github.com/steve-care-software/syntax/domain/identity/cryptography/hash"
	"github.com/steve-care-software/syntax/domain/identity/publics/assets"
)

type public struct {
	id          uuid.UUID
	name        string
	encryption  keys.PublicKey
	signature   hash.Hash
	host        string
	port        uint
	connections connections.Connections
	assets      assets.Assets
}

func createPublic(
	id uuid.UUID,
	name string,
	encryption keys.PublicKey,
	signature hash.Hash,
	host string,
	port uint,
) Public {
	return createPublicInternally(id, name, encryption, signature, host, port, nil, nil)
}

func createPublicWithConnections(
	id uuid.UUID,
	name string,
	encryption keys.PublicKey,
	signature hash.Hash,
	host string,
	port uint,
	connections connections.Connections,
) Public {
	return createPublicInternally(id, name, encryption, signature, host, port, connections, nil)
}

func createPublicWithAssets(
	id uuid.UUID,
	name string,
	encryption keys.PublicKey,
	signature hash.Hash,
	host string,
	port uint,
	assets assets.Assets,
) Public {
	return createPublicInternally(id, name, encryption, signature, host, port, nil, assets)
}

func createPublicWithConnectionsAndAssets(
	id uuid.UUID,
	name string,
	encryption keys.PublicKey,
	signature hash.Hash,
	host string,
	port uint,
	connections connections.Connections,
	assets assets.Assets,
) Public {
	return createPublicInternally(id, name, encryption, signature, host, port, connections, assets)
}

func createPublicInternally(
	id uuid.UUID,
	name string,
	encryption keys.PublicKey,
	signature hash.Hash,
	host string,
	port uint,
	connections connections.Connections,
	assets assets.Assets,
) Public {
	out := public{
		id:          id,
		name:        name,
		encryption:  encryption,
		signature:   signature,
		host:        host,
		port:        port,
		connections: connections,
		assets:      assets,
	}

	return &out
}

// ID returns the id
func (obj *public) ID() uuid.UUID {
	return obj.id
}

// Name returns the name
func (obj *public) Name() string {
	return obj.name
}

// Encryption returns the encryption
func (obj *public) Encryption() keys.PublicKey {
	return obj.encryption
}

// Encryption returns the encryption
func (obj *public) Signature() hash.Hash {
	return obj.signature
}

// Host returns the host
func (obj *public) Host() string {
	return obj.host
}

// Port returns the port
func (obj *public) Port() uint {
	return obj.port
}

// HasConnections returns true if there is connections, false otherwise
func (obj *public) HasConnections() bool {
	return obj.connections != nil
}

// Connections returns the connections, if any
func (obj *public) Connections() connections.Connections {
	return obj.connections
}

// HasAssets returns true if there is assets, false otherwise
func (obj *public) HasAssets() bool {
	return obj.assets != nil
}

// Assets returns the assets, if any
func (obj *public) Assets() assets.Assets {
	return obj.assets
}
