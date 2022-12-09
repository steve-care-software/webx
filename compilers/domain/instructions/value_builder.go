package instructions

import "errors"

type valueBuilder struct {
	variable     []byte
	constant     []byte
	instructions Instructions
	execution    []byte
}

func createValueBuilder() ValueBuilder {
	out := valueBuilder{
		variable:     nil,
		constant:     nil,
		instructions: nil,
		execution:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *valueBuilder) Create() ValueBuilder {
	return createValueBuilder()
}

// WithVariable adds a variable to the builder
func (app *valueBuilder) WithVariable(variable []byte) ValueBuilder {
	app.variable = variable
	return app
}

// WithConstant adds a constant to the builder
func (app *valueBuilder) WithConstant(constant []byte) ValueBuilder {
	app.constant = constant
	return app
}

// WithInstructions add instructions to the builder
func (app *valueBuilder) WithInstructions(instructions Instructions) ValueBuilder {
	app.instructions = instructions
	return app
}

// WithExecution add execution to the builder
func (app *valueBuilder) WithExecution(execution []byte) ValueBuilder {
	app.execution = execution
	return app
}

// Now builds a new Value instance
func (app *valueBuilder) Now() (Value, error) {
	if app.variable != nil {
		return createValueWithVariable(app.variable), nil
	}

	if app.constant != nil {
		return createValueWithConstant(app.constant), nil
	}

	if app.instructions != nil {
		return createValueWithInstructions(app.instructions), nil
	}

	if app.execution != nil {
		return createValueWithExecution(app.execution), nil
	}

	return nil, errors.New("the Value is invalid")
}
