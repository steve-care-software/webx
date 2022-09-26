package connections

import (
	"errors"

	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/syntax/domain/identity/cryptography/encryptions/keys"
)

type connectionBuilder struct {
	pID        *uuid.UUID
	pPublic    *uuid.UUID
	encryption keys.PrivateKey
}

func createConnectionBuilder() ConnectionBuilder {
	out := connectionBuilder{
		pID:        nil,
		pPublic:    nil,
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
func (app *connectionBuilder) WithPublic(public uuid.UUID) ConnectionBuilder {
	app.pPublic = &public
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

	if app.pPublic == nil {
		return nil, errors.New("the public ID is mandatory in order to build a Connection instance")
	}

	if app.encryption == nil {
		return nil, errors.New("the encryption is mandatory in order to build a Connection instance")
	}

	return createConnection(*app.pID, *app.pPublic, app.encryption), nil
}
