package layers

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/instances/executions/layers/results"
	source_layers "github.com/steve-care-software/datastencil/domain/instances/layers"
	"github.com/steve-care-software/historydb/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	input       []byte
	source      source_layers.Layer
	result      results.Result
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		input:       nil,
		source:      nil,
		result:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithInput adds an input to the builder
func (app *builder) WithInput(input []byte) Builder {
	app.input = input
	return app
}

// WithSource adds a source to the builder
func (app *builder) WithSource(source source_layers.Layer) Builder {
	app.source = source
	return app
}

// WithResult adds a result to the builder
func (app *builder) WithResult(result results.Result) Builder {
	app.result = result
	return app
}

// Now builds a new Layer instance
func (app *builder) Now() (Layer, error) {
	if app.input != nil && len(app.input) <= 0 {
		app.input = nil
	}

	if app.input == nil {
		return nil, errors.New("the input is mandatory in order to build a Layer instance")
	}

	if app.source == nil {
		return nil, errors.New("the source is mandatory in order to build a Layer instance")
	}

	if app.result == nil {
		return nil, errors.New("the result is mandatory in order to build a Layer instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.input,
		app.source.Hash().Bytes(),
		app.result.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createLayer(*pHash, app.input, app.source, app.result), nil
}
