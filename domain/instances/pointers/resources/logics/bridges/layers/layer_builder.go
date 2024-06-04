package layers

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/outputs"
)

type builder struct {
	hashAdapter  hash.Adapter
	input        string
	instructions instructions.Instructions
	output       outputs.Output
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter:  hashAdapter,
		input:        "",
		instructions: nil,
		output:       nil,
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
