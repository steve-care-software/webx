package decrypts

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	cipher      string
	account     string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		cipher:      "",
		account:     "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithCipher adds a cipher to the builder
func (app *builder) WithCipher(cipher string) Builder {
	app.cipher = cipher
	return app
}

// WithAccount adds an account to the builder
func (app *builder) WithAccount(account string) Builder {
	app.account = account
	return app
}

// Now bulds a new Decrypt instance
func (app *builder) Now() (Decrypt, error) {
	if app.cipher == "" {
		return nil, errors.New("the cipher is mandatory in order to build a Decrypt instance")
	}

	if app.account == "" {
		return nil, errors.New("the account is mandatory in order to build a Decrypt instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.cipher),
		[]byte(app.account),
	})

	if err != nil {
		return nil, err
	}

	return createDecrypt(*pHash, app.cipher, app.account), nil
}
