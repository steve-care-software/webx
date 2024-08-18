package executions

import "github.com/steve-care-software/webx/engine/domain/grammars/tokens"

type builder struct {
	tokens tokens.Tokens
	fnName string
}

func createBuilder() Builder {
	out := builder{
		tokens: nil,
		fnName: "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithTokens add tokens to the buiilder
func (app *builder) WithTokens(tokens tokens.Tokens) Builder {
	app.tokens = tokens
	return app
}

// WithFuncName add func name to the buiilder
func (app *builder) WithFuncName(funcName string) Builder {
	app.fnName = funcName
	return app
}

// Now builds a new Execution instance
func (app *builder) Now() (Execution, error) {
	if app.tokens != nil {
		return createExecutionWithTokens(app.fnName, app.tokens), nil
	}

	return createExecution(app.fnName), nil
}
