package encrypts

import (
	"errors"

	"github.com/steve-care-software/webx/engine/databases/entities/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	message     string
	password    string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		message:     "",
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

// WithMessage adds a message to the builder
func (app *builder) WithMessage(message string) Builder {
	app.message = message
	return app
}

// WithPassword adds a password to the builder
func (app *builder) WithPassword(password string) Builder {
	app.password = password
	return app
}

// Now builds a new Encrypt instance
func (app *builder) Now() (Encrypt, error) {
	if app.message == "" {
		return nil, errors.New("the message is mandatory in order to build a Encrypt instance")
	}

	if app.password == "" {
		return nil, errors.New("the password is mandatory in order to build a Encrypt instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.message),
		[]byte(app.password),
	})

	if err != nil {
		return nil, err
	}

	return createEncrypt(*pHash, app.message, app.password), nil
}
