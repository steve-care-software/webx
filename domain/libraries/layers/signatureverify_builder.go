package layers

import (
	"errors"

	"github.com/steve-care-software/identity/domain/hash"
)

type signatureVerifyBuilder struct {
	hashAdapter hash.Adapter
	signature   string
	message     string
}

func createSignatureVerifyBuilder(
	hashAdapter hash.Adapter,
) SignatureVerifyBuilder {
	out := signatureVerifyBuilder{
		hashAdapter: hashAdapter,
		signature:   "",
		message:     "",
	}

	return &out
}

// Create initializes the builder
func (app *signatureVerifyBuilder) Create() SignatureVerifyBuilder {
	return createSignatureVerifyBuilder(
		app.hashAdapter,
	)
}

// WithSignature adds a signature to the builder
func (app *signatureVerifyBuilder) WithSignature(signature string) SignatureVerifyBuilder {
	app.signature = signature
	return app
}

// WithMessage adds a message to the builder
func (app *signatureVerifyBuilder) WithMessage(message string) SignatureVerifyBuilder {
	app.message = message
	return app
}

// Now builds a new SignatureVerify instance
func (app *signatureVerifyBuilder) Now() (SignatureVerify, error) {
	if app.signature == "" {
		return nil, errors.New("the signature variable is mandatory in order to build a SignatureVerify instance")
	}

	if app.message == "" {
		return nil, errors.New("the message is mandatory in order to build a SignatureVerify instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.signature),
		[]byte(app.message),
	})

	if err != nil {
		return nil, err
	}

	return createSignatureVerify(*pHash, app.signature, app.message), nil
}
