package keys

import (
	"errors"
	"time"

	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/singles/keys/encryptors"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/singles/keys/profiles"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/singles/keys/signers"
)

type keyBuilder struct {
	profile    profiles.Profile
	encryptor  encryptors.Encryptor
	signer     signers.Signer
	pCreatedOn *time.Time
}

func createKeyBuilder() KeyBuilder {
	out := keyBuilder{
		profile:    nil,
		encryptor:  nil,
		signer:     nil,
		pCreatedOn: nil,
	}

	return &out
}

// Create initializes the builder
func (app *keyBuilder) Create() KeyBuilder {
	return createKeyBuilder()
}

// WithProfile adds a profile to the builder
func (app *keyBuilder) WithProfile(profile profiles.Profile) KeyBuilder {
	app.profile = profile
	return app
}

// WithEncryptor adds an encryptor to the builder
func (app *keyBuilder) WithEncryptor(encryptor encryptors.Encryptor) KeyBuilder {
	app.encryptor = encryptor
	return app
}

// WithSigner adds a signer to the builder
func (app *keyBuilder) WithSigner(signer signers.Signer) KeyBuilder {
	app.signer = signer
	return app
}

// CreatedOn adds a creation time to the builder
func (app *keyBuilder) CreatedOn(createdOn time.Time) KeyBuilder {
	app.pCreatedOn = &createdOn
	return app
}

// Now builds a new Key instance
func (app *keyBuilder) Now() (Key, error) {
	if app.profile == nil {
		return nil, errors.New("Tthe profile is mandatory in order to build a Key instance")
	}

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
		app.profile,
		app.encryptor,
		app.signer,
		*app.pCreatedOn,
	), nil
}
