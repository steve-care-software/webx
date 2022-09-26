package publics

import (
	"errors"

	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/syntax/domain/identity/cryptography/encryptions/keys"
	"github.com/steve-care-software/syntax/domain/identity/cryptography/hash"
)

type publicBuilder struct {
	pID        *uuid.UUID
	name       string
	encryption keys.PublicKey
	signature  hash.Hash
	host       string
	pPort      *uint
}

func createPublicBuilder() PublicBuilder {
	out := publicBuilder{
		pID:        nil,
		name:       "",
		encryption: nil,
		signature:  nil,
		host:       "",
		pPort:      nil,
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

	return createPublic(*app.pID, app.name, app.encryption, app.signature, app.host, *app.pPort), nil
}
