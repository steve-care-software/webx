package layers

import (
	"errors"

	"github.com/steve-care-software/identity/domain/hash"
)

type identityBuilder struct {
	hashAdapter hash.Adapter
	signer      Signer
	encryptor   Encryptor
}

func createIdentityBuilder(
	hashAdapter hash.Adapter,
) IdentityBuilder {
	out := identityBuilder{
		hashAdapter: hashAdapter,
		signer:      nil,
		encryptor:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *identityBuilder) Create() IdentityBuilder {
	return createIdentityBuilder(
		app.hashAdapter,
	)
}

// WithSigner adds a signer to the builder
func (app *identityBuilder) WithSigner(signer Signer) IdentityBuilder {
	app.signer = signer
	return app
}

// WithEncryptor adds an encryptor to the builder
func (app *identityBuilder) WithEncryptor(encryptor Encryptor) IdentityBuilder {
	app.encryptor = encryptor
	return app
}

// Now builds a new Identity instance
func (app *identityBuilder) Now() (Identity, error) {
	data := [][]byte{}
	if app.signer != nil {
		data = append(data, []byte("signer"))
		data = append(data, app.signer.Hash().Bytes())
	}

	if app.encryptor != nil {
		data = append(data, []byte("encryptor"))
		data = append(data, app.encryptor.Hash().Bytes())
	}

	if len(data) <= 0 {
		return nil, errors.New("the Identity is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.signer != nil {
		return createIdentityWithSigner(*pHash, app.signer), nil
	}

	return createIdentityWithEncryptor(*pHash, app.encryptor), nil
}
