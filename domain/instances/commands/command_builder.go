package commands

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/commands/results"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers"
)

type commandBuilder struct {
	hashAdapter hash.Adapter
	input       []byte
	layer       layers.Layer
	result      results.Result
	parent      Link
}

func createCommandBuilder(
	hashAdapter hash.Adapter,
) CommandBuilder {
	out := commandBuilder{
		hashAdapter: hashAdapter,
		input:       nil,
		layer:       nil,
		result:      nil,
		parent:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *commandBuilder) Create() CommandBuilder {
	return createCommandBuilder(
		app.hashAdapter,
	)
}

// WithInput adds an input to the builder
func (app *commandBuilder) WithInput(input []byte) CommandBuilder {
	app.input = input
	return app
}

// WithLayer adds a layer to the builder
func (app *commandBuilder) WithLayer(layer layers.Layer) CommandBuilder {
	app.layer = layer
	return app
}

// WithResult adds a result to the builder
func (app *commandBuilder) WithResult(result results.Result) CommandBuilder {
	app.result = result
	return app
}

// WithParent adds a parent to the builder
func (app *commandBuilder) WithParent(parent Link) CommandBuilder {
	app.parent = parent
	return app
}

// Now builds a new Command instance
func (app *commandBuilder) Now() (Command, error) {
	if app.input != nil && len(app.input) <= 0 {
		app.input = nil
	}

	if app.input == nil {
		return nil, errors.New("the input is mandatory in order to build a Command instance")
	}

	if app.layer == nil {
		return nil, errors.New("the layer is mandatory in order to build a Command instance")
	}

	if app.result == nil {
		return nil, errors.New("the result is mandatory in order to build a Command instance")
	}

	if app.parent == nil {
		return nil, errors.New("the parent is mandatory in order to build a Command instance")
	}

	data := [][]byte{
		app.input,
		app.layer.Hash().Bytes(),
		app.result.Hash().Bytes(),
		app.parent.Hash().Bytes(),
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createCommand(*pHash, app.input, app.layer, app.result, app.parent), nil
}
