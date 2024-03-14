package results

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/outputs/kinds"
)

type successBuilder struct {
	hashAdapter hash.Adapter
	output      Output
	kind        kinds.Kind
}

func createSuccessBuilder(
	hashAdapter hash.Adapter,
) SuccessBuilder {
	out := successBuilder{
		hashAdapter: hashAdapter,
		output:      nil,
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

// WithOutput adds an output to the builder
func (app *successBuilder) WithOutput(output Output) SuccessBuilder {
	app.output = output
	return app
}

// WithKind add kind to the builder
func (app *successBuilder) WithKind(kind kinds.Kind) SuccessBuilder {
	app.kind = kind
	return app
}

// Now builds a new Success instance
func (app *successBuilder) Now() (Success, error) {
	if app.output == nil {
		return nil, errors.New("the output is mandatory in order to build a Success instance")
	}

	if app.kind == nil {
		return nil, errors.New("the kind is mandatory in order to build a Success instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.output.Hash().Bytes(),
		app.kind.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createSuccess(*pHash, app.output, app.kind), nil
}
