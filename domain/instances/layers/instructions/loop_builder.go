package instructions

import (
	"errors"

	"github.com/steve-care-software/historydb/domain/hash"
)

type loopBuilder struct {
	hashAdapter  hash.Adapter
	amount       string
	instructions Instructions
}

func createLoopBuilder(
	hashAdapter hash.Adapter,
) LoopBuilder {
	out := loopBuilder{
		hashAdapter:  hashAdapter,
		amount:       "",
		instructions: nil,
	}

	return &out
}

// Create initializes the builder
func (app *loopBuilder) Create() LoopBuilder {
	return createLoopBuilder(
		app.hashAdapter,
	)
}

// WithAmount adds an amount to the builder
func (app *loopBuilder) WithAmount(amount string) LoopBuilder {
	app.amount = amount
	return app
}

// WithInstructions add instructions to the builder
func (app *loopBuilder) WithInstructions(instructions Instructions) LoopBuilder {
	app.instructions = instructions
	return app
}

// Now builds a new Loop instance
func (app *loopBuilder) Now() (Loop, error) {
	if app.amount == "" {
		return nil, errors.New("the amount is mandatory in order to build a Loop instance")
	}

	if app.instructions == nil {
		return nil, errors.New("the instructions is mandatory in order to build a Loop instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.amount),
		app.instructions.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createLoop(*pHash, app.amount, app.instructions), nil
}
