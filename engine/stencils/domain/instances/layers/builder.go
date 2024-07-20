package layers

import (
	"errors"

	"github.com/steve-care-software/webx/engine/states/domain/hash"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/outputs"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/references"
)

type builder struct {
	hashAdapter  hash.Adapter
	input        string
	instructions instructions.Instructions
	output       outputs.Output
	references   references.References
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter:  hashAdapter,
		input:        "",
		instructions: nil,
		output:       nil,
		references:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithInstructions add instructions to the builder
func (app *builder) WithInstructions(instructions instructions.Instructions) Builder {
	app.instructions = instructions
	return app
}

// WithOutput add output to the builder
func (app *builder) WithOutput(output outputs.Output) Builder {
	app.output = output
	return app
}

// WithInput adds an input to the builder
func (app *builder) WithInput(input string) Builder {
	app.input = input
	return app
}

// WithReferences add references to the builder
func (app *builder) WithReferences(references references.References) Builder {
	app.references = references
	return app
}

// Now builds a new Layer instance
func (app *builder) Now() (Layer, error) {
	if app.instructions == nil {
		return nil, errors.New("the instructions is mandatory in order to build a Layer instance")
	}

	if app.output == nil {
		return nil, errors.New("the output is mandatory in order to build a Layer instance")
	}

	if app.input == "" {
		return nil, errors.New("the input is mandatory in order to build a Layer instance")
	}

	data := [][]byte{
		app.instructions.Hash().Bytes(),
		app.output.Hash().Bytes(),
		[]byte(app.input),
	}

	if app.references != nil {
		data = append(data, app.references.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.references != nil {
		return createLayerWithReferences(*pHash, app.instructions, app.output, app.input, app.references), nil
	}

	return createLayer(*pHash, app.instructions, app.output, app.input), nil

}
