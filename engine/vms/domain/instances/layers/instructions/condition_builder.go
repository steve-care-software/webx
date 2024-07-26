package instructions

import (
	"errors"

	"github.com/steve-care-software/webx/engine/databases/entities/domain/hash"
)

type conditionBuilder struct {
	hashAdapter  hash.Adapter
	variable     string
	instructions Instructions
}

func createConditionBuilder(
	hashAdapter hash.Adapter,
) ConditionBuilder {
	out := conditionBuilder{
		hashAdapter:  hashAdapter,
		variable:     "",
		instructions: nil,
	}

	return &out
}

// Create initializes the builder
func (app *conditionBuilder) Create() ConditionBuilder {
	return createConditionBuilder(
		app.hashAdapter,
	)
}

// WithVariable adds a variable to the builder
func (app *conditionBuilder) WithVariable(variable string) ConditionBuilder {
	app.variable = variable
	return app
}

// WithInstructions add instructions to the builder
func (app *conditionBuilder) WithInstructions(instructions Instructions) ConditionBuilder {
	app.instructions = instructions
	return app
}

// Now builds a new Condition instance
func (app *conditionBuilder) Now() (Condition, error) {
	if app.variable == "" {
		return nil, errors.New("the variable is mandatory in order to build a Condition instance")
	}

	if app.instructions == nil {
		return nil, errors.New("the instructions is mandatory in order to build a Condition instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.variable),
		app.instructions.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createCondition(*pHash, app.variable, app.instructions), nil
}
