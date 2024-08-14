package encryptors

import (
	"crypto/rsa"
	"errors"
)

type publicKeyBuilder struct {
	pKey *rsa.PublicKey
}

func createPublicKeyBuilder() PublicKeyBuilder {
	out := publicKeyBuilder{
		pKey: nil,
	}

	return &out
}

// Create initializes the builder
func (app *publicKeyBuilder) Create() PublicKeyBuilder {
	return createPublicKeyBuilder()
}

// WithKey adds a key to the builder
func (app *publicKeyBuilder) WithKey(key rsa.PublicKey) PublicKeyBuilder {
	app.pKey = &key
	return app
}

// Now builds a new PublicKey instance
func (app *publicKeyBuilder) Now() (PublicKey, error) {
	if app.pKey == nil {
		return nil, errors.New("the key is mandatory in order to build a PublicKey instance")
	}

	return createPublicKey(*app.pKey), nil
}
