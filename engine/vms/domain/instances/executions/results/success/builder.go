package success

import (
	"errors"

	"github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/executions/results/success/outputs"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/outputs/kinds"
)

type builder struct {
	hashAdapter hash.Adapter
	output      outputs.Output
	kind        kinds.Kind
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		output:      nil,
		kind:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithOutput adds an output to the builder
func (app *builder) WithOutput(output outputs.Output) Builder {
	app.output = output
	return app
}

// WithKind add kind to the builder
func (app *builder) WithKind(kind kinds.Kind) Builder {
	app.kind = kind
	return app
}

// Now builds a new Success instance
func (app *builder) Now() (Success, error) {
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
