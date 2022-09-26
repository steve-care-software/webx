package publics

import (
	"errors"

	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/syntax/domain/identity/connections"
	"github.com/steve-care-software/syntax/domain/identity/cryptography/encryptions/keys"
	"github.com/steve-care-software/syntax/domain/identity/cryptography/hash"
	"github.com/steve-care-software/syntax/domain/identity/publics/assets"
)

type publicBuilder struct {
	pID         *uuid.UUID
	name        string
	encryption  keys.PublicKey
	signature   hash.Hash
	host        string
	pPort       *uint
	connections connections.Connections
	assets      assets.Assets
}

func createPublicBuilder() PublicBuilder {
	out := publicBuilder{
		pID:         nil,
		name:        "",
		encryption:  nil,
		signature:   nil,
		host:        "",
		pPort:       nil,
		connections: nil,
		assets:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *publicBuilder) Create() PublicBuilder {
	return createPublicBuilder()
}

// WithID adds an ID to the builder
func (app *publicBuilder) WithID(id uuid.UUID) PublicBuilder {
	app.pID = &id
	return app
}

// WithName adds a name to the builder
func (app *publicBuilder) WithName(name string) PublicBuilder {
	app.name = name
	return app
}

// WithEncryption adds an encryption to the builder
func (app *publicBuilder) WithEncryption(encryption keys.PublicKey) PublicBuilder {
	app.encryption = encryption
	return app
}

// WithSignature adds a signature to the builder
func (app *publicBuilder) WithSignature(signature hash.Hash) PublicBuilder {
	app.signature = signature
	return app
}

// WithHost adds a host to the builder
func (app *publicBuilder) WithHost(host string) PublicBuilder {
	app.host = host
	return app
}

// WithPort adds a port to the builder
func (app *publicBuilder) WithPort(port uint) PublicBuilder {
	app.pPort = &port
	return app
}

// WithConnections adds a connection to the builder
func (app *publicBuilder) WithConnections(connections connections.Connections) PublicBuilder {
	app.connections = connections
	return app
}

// WithAssets add assets to the builder
func (app *publicBuilder) WithAssets(assets assets.Assets) PublicBuilder {
	app.assets = assets
	return app
}

// Now builds a new Public instance
func (app *publicBuilder) Now() (Public, error) {
	if app.pID == nil {
		return nil, errors.New("the ID is mandatory in order to build a Public instance")
	}

	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Public instance")
	}

	if app.encryption == nil {
		return nil, errors.New("the encryption is mandatory in order to build a Public instance")
	}

	if app.signature == nil {
		return nil, errors.New("the signature is mandatory in order to build a Public instance")
	}

	if app.host == "" {
		return nil, errors.New("the host is mandatory in order to build a Public instance")
	}

	if app.pPort == nil {
		return nil, errors.New("the port is mandatory in order to build a Public instance")
	}

	if app.connections != nil && app.assets != nil {
		return createPublicWithConnectionsAndAssets(*app.pID, app.name, app.encryption, app.signature, app.host, *app.pPort, app.connections, app.assets), nil
	}

	if app.connections != nil {
		return createPublicWithConnections(*app.pID, app.name, app.encryption, app.signature, app.host, *app.pPort, app.connections), nil
	}

	if app.assets != nil {
		return createPublicWithAssets(*app.pID, app.name, app.encryption, app.signature, app.host, *app.pPort, app.assets), nil
	}

	return createPublic(*app.pID, app.name, app.encryption, app.signature, app.host, *app.pPort), nil
}
