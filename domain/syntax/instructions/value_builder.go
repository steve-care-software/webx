package instructions

import "errors"

type valueBuilder struct {
	input        string
	str          string
	instructions Instructions
	execution    string
}

func createValueBuilder() ValueBuilder {
	out := valueBuilder{
		input:        "",
		str:          "",
		instructions: nil,
		execution:    "",
	}

	return &out
}

// Create initializes the builder
func (app *valueBuilder) Create() ValueBuilder {
	return createValueBuilder()
}

// WithInput adds an input to the builder
func (app *valueBuilder) WithInput(input string) ValueBuilder {
	app.input = input
	return app
}

// WithString adds a string to the builder
func (app *valueBuilder) WithString(str string) ValueBuilder {
	app.str = str
	return app
}

// WithInstructions add instructions to the builder
func (app *valueBuilder) WithInstructions(instructions Instructions) ValueBuilder {
	app.instructions = instructions
	return app
}

// WithExecution add execution to the builder
func (app *valueBuilder) WithExecution(execution string) ValueBuilder {
	app.execution = execution
	return app
}

// Now builds a new Value instance
func (app *valueBuilder) Now() (Value, error) {
	if app.input != "" {
		return createValueWithInput(app.input), nil
	}

	if app.str != "" {
		return createValueWithString(app.str), nil
	}

	if app.instructions != nil {
		return createValueWithInstructions(app.instructions), nil
	}

	if app.execution != "" {
		return createValueWithExecution(app.execution), nil
	}

	return nil, errors.New("the Value is invalid")
}
