package encrypts

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	message     string
	account     string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		message:     "",
		account:     "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithMessage adds a message to the builder
func (app *builder) WithMessage(message string) Builder {
	app.message = message
	return app
}

// WithAccount adds an account to the builder
func (app *builder) WithAccount(account string) Builder {
	app.account = account
	return app
}

// Now bulds a new Encrypt instance
func (app *builder) Now() (Encrypt, error) {
	if app.message == "" {
		return nil, errors.New("the message is mandatory in order to build a Encrypt instance")
	}

	if app.account == "" {
		return nil, errors.New("the account is mandatory in order to build a Encrypt instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.message),
		[]byte(app.account),
	})

	if err != nil {
		return nil, err
	}

	return createEncrypt(*pHash, app.message, app.account), nil
}
