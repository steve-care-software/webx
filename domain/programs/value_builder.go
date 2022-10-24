package programs

import "errors"

type valueBuilder struct {
	input     string
	str       string
	execution Application
	program   Program
}

func createValueBuilder() ValueBuilder {
	out := valueBuilder{
		input:     "",
		str:       "",
		execution: nil,
		program:   nil,
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

// WithExecution adds an execution to the builder
func (app *valueBuilder) WithExecution(execution Application) ValueBuilder {
	app.execution = execution
	return app
}

// WithProgram adds a program to the builder
func (app *valueBuilder) WithProgram(program Program) ValueBuilder {
	app.program = program
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

	if app.execution != nil {
		return createValueWithExecution(app.execution), nil
	}

	if app.program != nil {
		return createValueWithProgram(app.program), nil
	}

	return nil, errors.New("the Value is invalid")
}
