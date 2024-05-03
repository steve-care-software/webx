package encrypts

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	message     string
	pubKey      string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		message:     "",
		pubKey:      "",
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
func (app *builder) WithMessage(msg string) Builder {
	app.message = msg
	return app
}

// WithPublicKey adds a publicKey to the builder
func (app *builder) WithPublicKey(pubKey string) Builder {
	app.pubKey = pubKey
	return app
}

// Now builds a new Encrypt instance
func (app *builder) Now() (Encrypt, error) {
	if app.message == "" {
		return nil, errors.New("the message is mandatory in order to build an Encrypt instance")
	}

	if app.pubKey == "" {
		return nil, errors.New("the publicKey is mandatory in order to build an Encrypt instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.message),
		[]byte(app.pubKey),
	})

	if err != nil {
		return nil, err
	}

	return createEncrypt(*pHash, app.message, app.pubKey), nil
}
