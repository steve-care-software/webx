package commits

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/accounts/signers"
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/commits/contents"
)

type builder struct {
	hashAdapter hash.Adapter
	content     contents.Content
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
func (app *builder) WithContent(content contents.Content) Builder {
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

	sigBytes, err := app.signature.Bytes()
	if err != nil {
		return nil, err
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.content.Hash().Bytes(),
		sigBytes,
	})

	if err != nil {
		return nil, err
	}

	return createCommit(*pHash, app.content, app.signature), nil
}
