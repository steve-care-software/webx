package layers

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type conditionResourceBuilder struct {
	hashAdapter  hash.Adapter
	variable     string
	instructions Instructions
}

func createConditionResourceBuilder(
	hashAdapter hash.Adapter,
) ConditionResourceBuilder {
	out := conditionResourceBuilder{
		hashAdapter:  hashAdapter,
		variable:     "",
		instructions: nil,
	}

	return &out
}

// Create initializes the builder
func (app *conditionResourceBuilder) Create() ConditionResourceBuilder {
	return createConditionResourceBuilder(
		app.hashAdapter,
	)
}

// WithVariable adds a variable to the builder
func (app *conditionResourceBuilder) WithVariable(variable string) ConditionResourceBuilder {
	app.variable = variable
	return app
}

// WithInstructions add instructions to the builder
func (app *conditionResourceBuilder) WithInstructions(instructions Instructions) ConditionResourceBuilder {
	app.instructions = instructions
	return app
}

// Now builds a new Condition instance
func (app *conditionResourceBuilder) Now() (Condition, error) {
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
