package connections

import (
	"errors"

	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/syntax/domain/identity/cryptography/encryptions/keys"
	"github.com/steve-care-software/syntax/domain/identity/publics"
)

type connectionBuilder struct {
	pID        *uuid.UUID
	public     publics.Public
	encryption keys.PrivateKey
}

func createConnectionBuilder() ConnectionBuilder {
	out := connectionBuilder{
		pID:        nil,
		public:     nil,
		encryption: nil,
	}

	return &out
}

// Create initializes the builder
func (app *connectionBuilder) Create() ConnectionBuilder {
	return createConnectionBuilder()
}

// WithID adds an id to the builder
func (app *connectionBuilder) WithID(id uuid.UUID) ConnectionBuilder {
	app.pID = &id
	return app
}

// WithPublic adds a public to the builder
func (app *connectionBuilder) WithPublic(public publics.Public) ConnectionBuilder {
	app.public = public
	return app
}

// WithEncryption adds anencryption's private key to the builder
func (app *connectionBuilder) WithEncryption(encryption keys.PrivateKey) ConnectionBuilder {
	app.encryption = encryption
	return app
}

// Now builds a new Connection instance
func (app *connectionBuilder) Now() (Connection, error) {
	if app.pID == nil {
		return nil, errors.New("the ID is mandatory in order to build a Connection instance")
	}

	if app.public == nil {
		return nil, errors.New("the public is mandatory in order to build a Connection instance")
	}

	if app.encryption == nil {
		return nil, errors.New("the encryption is mandatory in order to build a Connection instance")
	}

	return createConnection(*app.pID, app.public, app.encryption), nil
}
