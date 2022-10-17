package identities

import (
	"errors"
	"time"

	"github.com/steve-care-software/syntax/domain/syntax/databases"
	"github.com/steve-care-software/syntax/domain/syntax/databases/cryptography/encryptions/keys"
	"github.com/steve-care-software/syntax/domain/syntax/databases/cryptography/signatures"
	"github.com/steve-care-software/syntax/domain/syntax/identities/modifications"
)

type builder struct {
	name          string
	sig           signatures.PrivateKey
	enc           keys.PrivateKey
	pCreatedOn    *time.Time
	databases     databases.Databases
	modifications modifications.Modifications
}

func createBuilder() Builder {
	out := builder{
		name:          "",
		sig:           nil,
		enc:           nil,
		pCreatedOn:    nil,
		databases:     nil,
		modifications: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithSignature adds a signature pk to the builder
func (app *builder) WithSignature(sig signatures.PrivateKey) Builder {
	app.sig = sig
	return app
}

// WithEncryption adds an encryption pk to the builder
func (app *builder) WithEncryption(enc keys.PrivateKey) Builder {
	app.enc = enc
	return app
}

// WithDatabases add databases to the builder
func (app *builder) WithDatabases(databases databases.Databases) Builder {
	app.databases = databases
	return app
}

// WithModifications add modifications to the builder
func (app *builder) WithModifications(modifications modifications.Modifications) Builder {
	app.modifications = modifications
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.pCreatedOn = &createdOn
	return app
}

// Now builds a new Identity instance
func (app *builder) Now() (Identity, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build an Identity instance")
	}

	if app.sig == nil {
		return nil, errors.New("the signature's private key is mandatory in order to build an Identity instance")
	}

	if app.enc == nil {
		return nil, errors.New("the encryption's private key is mandatory in order to build an Identity instance")
	}

	if app.pCreatedOn == nil {
		return nil, errors.New("the creation time is mandatory in order to build an Identity instance")
	}

	if app.databases != nil && app.modifications != nil {
		return createIdentityWithDatabasesAndModifications(app.name, app.sig, app.enc, *app.pCreatedOn, app.databases, app.modifications), nil
	}

	if app.databases != nil {
		return createIdentityWithDatabases(app.name, app.sig, app.enc, *app.pCreatedOn, app.databases), nil
	}

	if app.modifications != nil {
		return createIdentityWithModifications(app.name, app.sig, app.enc, *app.pCreatedOn, app.modifications), nil
	}

	return createIdentity(app.name, app.sig, app.enc, *app.pCreatedOn), nil
}
