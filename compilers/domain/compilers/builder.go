package compilers

import "errors"

type builder struct {
	outputs    []uint
	executions Executions
}

func createBuilder() Builder {
	out := builder{
		outputs:    nil,
		executions: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithOutputs adds an outputs to the builder
func (app *builder) WithOutputs(outputs []uint) Builder {
	app.outputs = outputs
	return app
}

// WithExecutions add executions to the builder
func (app *builder) WithExecutions(executions Executions) Builder {
	app.executions = executions
	return app
}

// Now builds a new Compiler instance
func (app *builder) Now() (Compiler, error) {
	if app.outputs != nil && len(app.outputs) <= 0 {
		app.outputs = nil
	}

	if app.executions == nil {
		return nil, errors.New("the executions is mandatory in order to build a Compiler instance")
	}

	if app.outputs != nil {
		return createCompilerWithOutputs(app.executions, app.outputs), nil
	}

	return createCompiler(app.executions), nil
}
