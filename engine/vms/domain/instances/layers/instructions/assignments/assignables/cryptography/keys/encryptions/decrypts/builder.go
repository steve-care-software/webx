package decrypts

import (
	"errors"

	"github.com/steve-care-software/webx/engine/databases/entities/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	cipher      string
	privateKey  string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		cipher:      "",
		privateKey:  "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithCipher adds a cipher to the builder
func (app *builder) WithCipher(cipher string) Builder {
	app.cipher = cipher
	return app
}

// WithPrivateKey adds a privateKey to the builder
func (app *builder) WithPrivateKey(privateKey string) Builder {
	app.privateKey = privateKey
	return app
}

// Now builds a new Decrypt instance
func (app *builder) Now() (Decrypt, error) {
	if app.cipher == "" {
		return nil, errors.New("the cipher is mandatory in order to build a Decrypt instance")
	}

	if app.privateKey == "" {
		return nil, errors.New("the privateKey is mandatory in order to build a Decrypt instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.cipher),
		[]byte(app.privateKey),
	})

	if err != nil {
		return nil, err
	}

	return createDecrypt(*pHash, app.cipher, app.privateKey), nil
}
