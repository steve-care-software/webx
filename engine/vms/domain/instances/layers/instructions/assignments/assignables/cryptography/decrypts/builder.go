package decrypts

import (
	"errors"

	"github.com/steve-care-software/webx/engine/hashes/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	cipher      string
	password    string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		cipher:      "",
		password:    "",
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

// WithPassword adds a password to the builder
func (app *builder) WithPassword(password string) Builder {
	app.password = password
	return app
}

// Now builds a new Decrypt instance
func (app *builder) Now() (Decrypt, error) {
	if app.cipher == "" {
		return nil, errors.New("the cipher is mandatory in order to build a Decrypt instance")
	}

	if app.password == "" {
		return nil, errors.New("the password is mandatory in order to build a Decrypt instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.cipher),
		[]byte(app.password),
	})

	if err != nil {
		return nil, err
	}

	return createDecrypt(*pHash, app.cipher, app.password), nil
}
