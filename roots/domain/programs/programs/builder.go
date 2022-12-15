package programs

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"
)

type builder struct {
	hashAdapter  hash.Adapter
	instructions Instructions
	outputs      []uint
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter:  hashAdapter,
		instructions: nil,
		outputs:      nil,
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
func (app *builder) WithInstructions(instructions Instructions) Builder {
	app.instructions = instructions
	return app
}

// WithOutputs add outputs to the builder
func (app *builder) WithOutputs(outputs []uint) Builder {
	app.outputs = outputs
	return app
}

// Now builds a new Program instance
func (app *builder) Now() (Program, error) {
	if app.instructions == nil {
		return nil, errors.New("the instructions is mandatory in order to build a Program instance")
	}

	if app.outputs != nil && len(app.outputs) <= 0 {
		app.outputs = nil
	}

	data := [][]byte{
		app.instructions.Hash().Bytes(),
	}

	if app.outputs != nil {
		for _, oneOutput := range app.outputs {
			data = append(data, []byte(fmt.Sprintf("%d", oneOutput)))
		}
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.outputs != nil {
		return createProgramWithOutputs(*pHash, app.instructions, app.outputs), nil
	}

	return createProgram(*pHash, app.instructions), nil
}
