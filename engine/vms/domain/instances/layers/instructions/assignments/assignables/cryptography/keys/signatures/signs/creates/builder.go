package creates

import (
	"errors"

	"github.com/steve-care-software/webx/engine/databases/entities/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	message     string
	privateKey  string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		message:     "",
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

// WithMessage adds a message to the builder
func (app *builder) WithMessage(message string) Builder {
	app.message = message
	return app
}

// WithPrivateKey adds a privateKey to the builder
func (app *builder) WithPrivateKey(privateKey string) Builder {
	app.privateKey = privateKey
	return app
}

// Now builds a new Create instance
func (app *builder) Now() (Create, error) {
	if app.message == "" {
		return nil, errors.New("the message is mandatory in order to build a Create instance")
	}

	if app.privateKey == "" {
		return nil, errors.New("the privateKey is mandatory in order to build a Create instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.message),
		[]byte(app.privateKey),
	})

	if err != nil {
		return nil, err
	}

	return createCreate(*pHash, app.message, app.privateKey), nil
}
