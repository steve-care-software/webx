package keys

import (
	"crypto/rsa"
	"errors"
)

type publicKeyBuilder struct {
	key *rsa.PublicKey
}

func createPublicKeyBuilder() PublicKeyBuilder {
	out := publicKeyBuilder{
		key: nil,
	}

	return &out
}

// Create initializes the builder
func (app *publicKeyBuilder) Create() PublicKeyBuilder {
	return createPublicKeyBuilder()
}

// WithKey adds a key to the builder
func (app *publicKeyBuilder) WithKey(key rsa.PublicKey) PublicKeyBuilder {
	app.key = &key
	return app
}

// Now builds a new key instance
func (app *publicKeyBuilder) Now() (PublicKey, error) {
	if app.key == nil {
		return nil, errors.New("the rsa PublicKey is mandatory in order to build a PublicKey instance")
	}

	return createPublicKey(*app.key), nil
}
