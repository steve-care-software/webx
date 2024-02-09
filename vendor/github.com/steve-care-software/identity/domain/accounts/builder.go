package accounts

import (
	"errors"

	"github.com/steve-care-software/identity/domain/accounts/encryptors"
	"github.com/steve-care-software/identity/domain/accounts/signers"
)

type builder struct {
	username  string
	encryptor encryptors.Encryptor
	signer    signers.Signer
}

func createBuilder() Builder {
	out := builder{
		username:  "",
		encryptor: nil,
		signer:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithUsername adds a username to the builder
func (app *builder) WithUsername(username string) Builder {
	app.username = username
	return app
}

// WithEncryptor adds an encryptor to the builder
func (app *builder) WithEncryptor(encryptor encryptors.Encryptor) Builder {
	app.encryptor = encryptor
	return app
}

// WithSigner adds a signer to the builder
func (app *builder) WithSigner(signer signers.Signer) Builder {
	app.signer = signer
	return app
}

// Now builds a new Account instance
func (app *builder) Now() (Account, error) {
	if app.username == "" {
		return nil, errors.New("the username is mandatory in order to build an Account instance")
	}

	if app.encryptor == nil {
		return nil, errors.New("the encryptor is mandatory in order to build an Account instance")
	}

	if app.signer == nil {
		return nil, errors.New("the signer is mandatory in order to build an Account instance")
	}

	return createAccount(
		app.username,
		app.encryptor,
		app.signer,
	), nil
}
