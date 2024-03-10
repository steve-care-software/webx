package criterias

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/accounts/encryptors"
	"github.com/steve-care-software/datastencil/domain/accounts/signers"
)

type builder struct {
	signer    signers.Signer
	encryptor encryptors.Encryptor
	username  string
	password  []byte
}

func createBuilder() Builder {
	out := builder{
		signer:    nil,
		encryptor: nil,
		username:  "",
		password:  nil,
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

// WithPassword adds a password to the builder
func (app *builder) WithPassword(password []byte) Builder {
	app.password = password
	return app
}

// WithSigner adds a signer to the builder
func (app *builder) WithSigner(signer signers.Signer) Builder {
	app.signer = signer
	return app
}

// WithEncryptor adds an encryptor to the builder
func (app *builder) WithEncryptor(encryptor encryptors.Encryptor) Builder {
	app.encryptor = encryptor
	return app
}

// Now builds a new Criteria instance
func (app *builder) Now() (Criteria, error) {
	if app.password != nil && len(app.password) <= 0 {
		app.password = nil
	}

	if app.username == "" && app.password == nil && app.signer == nil && app.encryptor == nil {
		return nil, errors.New("the Criteria is invalid")
	}

	return createCriteria(
		app.signer,
		app.encryptor,
		app.username,
		app.password,
	), nil
}
