package programs

import (
	"errors"

	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"
)

type builder struct {
	pHash        *hash.Hash
	instructions []hash.Hash
	outputs      []uint
}

func createBuilder() Builder {
	out := builder{
		pHash:        nil,
		instructions: nil,
		outputs:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithHash adds an hash to the builder
func (app *builder) WithHash(hash hash.Hash) Builder {
	app.pHash = &hash
	return app
}

// WithInstructions add instructions to the builder
func (app *builder) WithInstructions(instructions []hash.Hash) Builder {
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
	if app.pHash == nil {
		return nil, errors.New("the hash is mandatory in order to build a Program instance")
	}

	if app.instructions != nil && len(app.instructions) <= 0 {
		app.instructions = nil
	}

	if app.instructions == nil {
		return nil, errors.New("there must be at least 1 Instruction in order to build a Program instance")
	}

	if app.outputs != nil && len(app.outputs) <= 0 {
		app.outputs = nil
	}

	if app.outputs != nil {
		return createProgramWithOutputs(*app.pHash, app.instructions, app.outputs), nil
	}

	return createProgram(*app.pHash, app.instructions), nil
}
