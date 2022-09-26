package identities

import (
	"errors"

	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/syntax/domain/identity/cryptography/signatures"
	"github.com/steve-care-software/syntax/domain/identity/identities/assets"
	"github.com/steve-care-software/syntax/domain/identity/identities/connections"
	"github.com/steve-care-software/syntax/domain/identity/identities/publics"
)

type builder struct {
	pID         *uuid.UUID
	public      publics.Public
	pk          signatures.PrivateKey
	connections connections.Connections
	assets      assets.Assets
}

func createBuilder() Builder {
	out := builder{
		pID:         nil,
		public:      nil,
		pk:          nil,
		connections: nil,
		assets:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithID adds an ID to the builder
func (app *builder) WithID(id uuid.UUID) Builder {
	app.pID = &id
	return app
}

// WithPublic adds a public to the builder
func (app *builder) WithPublic(public publics.Public) Builder {
	app.public = public
	return app
}

// WithPrivateKey adds a pk to the builder
func (app *builder) WithPrivateKey(pk signatures.PrivateKey) Builder {
	app.pk = pk
	return app
}

// WithConnections add connections to the builder
func (app *builder) WithConnections(connections connections.Connections) Builder {
	app.connections = connections
	return app
}

// WithAssets add assets to the builder
func (app *builder) WithAssets(assets assets.Assets) Builder {
	app.assets = assets
	return app
}

// Now builds a new Identity instance
func (app *builder) Now() (Identity, error) {
	if app.pID == nil {
		return nil, errors.New("the ID is mandatory in order to build an Identity instance")
	}

	if app.public == nil {
		return nil, errors.New("the public is mandatory in order to build an Identity instance")
	}

	if app.pk == nil {
		return nil, errors.New("the pk is mandatory in order to build an Identity instance")
	}

	if app.connections != nil && app.assets != nil {
		return createIdentityWithConnectionsAndAssets(*app.pID, app.public, app.pk, app.connections, app.assets), nil
	}

	if app.connections != nil {
		return createIdentityWithConnections(*app.pID, app.public, app.pk, app.connections), nil
	}

	if app.assets != nil {
		return createIdentityWithAssets(*app.pID, app.public, app.pk, app.assets), nil
	}

	return createIdentity(*app.pID, app.public, app.pk), nil
}
