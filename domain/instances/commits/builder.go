package commits

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/keys/signers"
)

type builder struct {
	hashAdapter hash.Adapter
	content     Content
	signature   signers.Signature
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		content:     nil,
		signature:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithContent adds content to the builder
func (app *builder) WithContent(content Content) Builder {
	app.content = content
	return app
}

// WithSignature adds signature to the builder
func (app *builder) WithSignature(signature signers.Signature) Builder {
	app.signature = signature
	return app
}

// Now builds a new Commit instance
func (app *builder) Now() (Commit, error) {
	if app.content == nil {
		return nil, errors.New("the content is mandatory in order to build a Commit instance")
	}

	if app.signature == nil {
		return nil, errors.New("the signature is mandatory in order to build a Commit instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.content.Hash().Bytes(),
		[]byte(app.signature.String()),
	})

	if err != nil {
		return nil, err
	}

	return createCommit(*pHash, app.content, app.signature), nil
}
