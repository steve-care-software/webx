package executions

import "github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/executions/parameters"

type builder struct {
	parameters parameters.Parameters
	fnName     string
}

func createBuilder() Builder {
	out := builder{
		parameters: nil,
		fnName:     "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithParameters add parameters to the buiilder
func (app *builder) WithParameters(parameters parameters.Parameters) Builder {
	app.parameters = parameters
	return app
}

// WithFuncName add func name to the buiilder
func (app *builder) WithFuncName(funcName string) Builder {
	app.fnName = funcName
	return app
}

// Now builds a new Execution instance
func (app *builder) Now() (Execution, error) {
	if app.parameters != nil {
		return createExecutionWithParameters(app.fnName, app.parameters), nil
	}

	return createExecution(app.fnName), nil
}
