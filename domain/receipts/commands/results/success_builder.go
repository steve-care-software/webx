package results

import (
	"errors"

	"github.com/steve-care-software/identity/domain/hash"
	"github.com/steve-care-software/datastencil/domain/libraries/layers"
)

type successBuilder struct {
	hashAdapter hash.Adapter
	bytes       []byte
	kind        layers.Kind
}

func createSuccessBuilder(
	hashAdapter hash.Adapter,
) SuccessBuilder {
	out := successBuilder{
		hashAdapter: hashAdapter,
		bytes:       nil,
		kind:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *successBuilder) Create() SuccessBuilder {
	return createSuccessBuilder(
		app.hashAdapter,
	)
}

// WithBytes add bytes to the builder
func (app *successBuilder) WithBytes(bytes []byte) SuccessBuilder {
	app.bytes = bytes
	return app
}

// WithKind add kind to the builder
func (app *successBuilder) WithKind(kind layers.Kind) SuccessBuilder {
	app.kind = kind
	return app
}

// Now builds a new Success instance
func (app *successBuilder) Now() (Success, error) {
	if app.bytes != nil && len(app.bytes) <= 0 {
		app.bytes = nil
	}

	if app.bytes == nil {
		return nil, errors.New("the bytes are mandatory in order to build a Success instance")
	}

	if app.kind == nil {
		return nil, errors.New("the kind is mandatory in order to build a Success instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.bytes,
		app.kind.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createSuccess(*pHash, app.bytes, app.kind), nil
}
