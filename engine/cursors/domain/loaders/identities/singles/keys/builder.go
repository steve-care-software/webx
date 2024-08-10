package keys

import (
	"errors"
	"time"

	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/singles/keys/encryptors"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/singles/keys/signers"
)

type builder struct {
	encryptor  encryptors.Encryptor
	signer     signers.Signer
	pCreatedOn *time.Time
}

func createBuilder() Builder {
	out := builder{
		encryptor:  nil,
		signer:     nil,
		pCreatedOn: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
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

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.pCreatedOn = &createdOn
	return app
}

// Now builds a new Key instance
func (app *builder) Now() (Key, error) {
	if app.encryptor == nil {
		return nil, errors.New("Tthe encryptor is mandatory in order to build a Key instance")
	}

	if app.signer == nil {
		return nil, errors.New("Tthe signer is mandatory in order to build a Key instance")
	}

	if app.pCreatedOn == nil {
		return nil, errors.New("Tthe creation time is mandatory in order to build a Key instance")
	}

	return createKey(
		app.encryptor,
		app.signer,
		*app.pCreatedOn,
	), nil
}
