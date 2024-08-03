package validates

import (
	"errors"

	"github.com/steve-care-software/webx/engine/hashes/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	signature   string
	message     string
	publicKey   string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		signature:   "",
		message:     "",
		publicKey:   "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithSignature adds a signature to the builder
func (app *builder) WithSignature(sig string) Builder {
	app.signature = sig
	return app
}

// WithMessage adds a message to the builder
func (app *builder) WithMessage(msg string) Builder {
	app.message = msg
	return app
}

// WithPublicKey adds a publicKey to the builder
func (app *builder) WithPublicKey(pubKey string) Builder {
	app.publicKey = pubKey
	return app
}

// Now builds a new Validate instance
func (app *builder) Now() (Validate, error) {
	if app.signature == "" {
		return nil, errors.New("the signature is mandatory in order to build a Validate instance")
	}

	if app.message == "" {
		return nil, errors.New("the message is mandatory in order to build a Validate instance")
	}

	if app.publicKey == "" {
		return nil, errors.New("the publicKey is mandatory in order to build a Validate instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.signature),
		[]byte(app.message),
		[]byte(app.publicKey),
	})

	if err != nil {
		return nil, err
	}

	return createValidate(*pHash, app.signature, app.message, app.publicKey), nil
}
