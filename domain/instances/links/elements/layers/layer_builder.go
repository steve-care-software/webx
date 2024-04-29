package layers

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/outputs"
)

type layerBuilder struct {
	hashAdapter  hash.Adapter
	input        string
	instructions instructions.Instructions
	output       outputs.Output
}

func createLayerBuilder(
	hashAdapter hash.Adapter,
) LayerBuilder {
	out := layerBuilder{
		hashAdapter:  hashAdapter,
		input:        "",
		instructions: nil,
		output:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *layerBuilder) Create() LayerBuilder {
	return createLayerBuilder(
		app.hashAdapter,
	)
}

// WithInstructions add instructions to the builder
func (app *layerBuilder) WithInstructions(instructions instructions.Instructions) LayerBuilder {
	app.instructions = instructions
	return app
}

// WithOutput add output to the builder
func (app *layerBuilder) WithOutput(output outputs.Output) LayerBuilder {
	app.output = output
	return app
}

// WithInput adds an input to the builder
func (app *layerBuilder) WithInput(input string) LayerBuilder {
	app.input = input
	return app
}

// Now builds a new Layer instance
func (app *layerBuilder) Now() (Layer, error) {
	if app.instructions == nil {
		return nil, errors.New("the instructions is mandatory in order to build a Layer instance")
	}

	if app.output == nil {
		return nil, errors.New("the output is mandatory in order to build a Layer instance")
	}

	if app.input == "" {
		return nil, errors.New("the input is mandatory in order to build a Layer instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.instructions.Hash().Bytes(),
		app.output.Hash().Bytes(),
		[]byte(app.input),
	})

	if err != nil {
		return nil, err
	}

	return createLayer(*pHash, app.instructions, app.output, app.input), nil

}
