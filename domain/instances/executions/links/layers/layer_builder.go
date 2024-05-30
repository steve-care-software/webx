package layers

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/executions/links/layers/results"
	source_layers "github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers"
)

type layerBuilder struct {
	hashAdapter hash.Adapter
	input       []byte
	source      source_layers.Layer
	result      results.Result
}

func createLayerBuilder(
	hashAdapter hash.Adapter,
) LayerBuilder {
	out := layerBuilder{
		hashAdapter: hashAdapter,
		input:       nil,
		source:      nil,
		result:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *layerBuilder) Create() LayerBuilder {
	return createLayerBuilder(
		app.hashAdapter,
	)
}

// WithInput adds an input to the builder
func (app *layerBuilder) WithInput(input []byte) LayerBuilder {
	app.input = input
	return app
}

// WithSource adds a source to the builder
func (app *layerBuilder) WithSource(source source_layers.Layer) LayerBuilder {
	app.source = source
	return app
}

// WithResult adds a result to the builder
func (app *layerBuilder) WithResult(result results.Result) LayerBuilder {
	app.result = result
	return app
}

// Now builds a new Layer instance
func (app *layerBuilder) Now() (Layer, error) {
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
