package encryptors

import (
	"crypto/rand"
	"crypto/rsa"
	"errors"
)

type builder struct {
	pubKeyBuilder PublicKeyBuilder
	pk            *rsa.PrivateKey
	bitRate       int
}

func createBuilder(
	pubKeyBuilder PublicKeyBuilder,
) Builder {
	out := builder{
		pubKeyBuilder: pubKeyBuilder,
		pk:            nil,
		bitRate:       0,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.pubKeyBuilder)
}

// WithPK adds a privateKey to the builder
func (app *builder) WithPK(pk rsa.PrivateKey) Builder {
	app.pk = &pk
	return app
}

// WithBitRate adds a bitRate to the builder
func (app *builder) WithBitRate(bitRate int) Builder {
	app.bitRate = bitRate
	return app
}

// Now builds a new Encryptor instance
func (app *builder) Now() (Encryptor, error) {
	if app.bitRate > 0 {
		pk, err := rsa.GenerateKey(rand.Reader, app.bitRate)
		if err != nil {
			return nil, err
		}

		app.pk = pk
	}

	if app.pk == nil {
		return nil, errors.New("the rsa PrivateKey is mandatory in order to build an Encryptor instance")
	}

	pubKey, err := app.pubKeyBuilder.Create().
		WithKey(app.pk.PublicKey).
		Now()

	if err != nil {
		return nil, err
	}

	return createEncryptor(*app.pk, pubKey), nil
}
