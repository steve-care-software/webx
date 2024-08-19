package executions

import "github.com/steve-care-software/webx/engine/domain/grammars/blocks/lines/tokens/elements"

type builder struct {
	elements elements.Elements
	fnName   string
}

func createBuilder() Builder {
	out := builder{
		elements: nil,
		fnName:   "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithElements add elements to the buiilder
func (app *builder) WithElements(elements elements.Elements) Builder {
	app.elements = elements
	return app
}

// WithFuncName add func name to the buiilder
func (app *builder) WithFuncName(funcName string) Builder {
	app.fnName = funcName
	return app
}

// Now builds a new Execution instance
func (app *builder) Now() (Execution, error) {
	if app.elements != nil {
		return createExecutionWithElements(app.fnName, app.elements), nil
	}

	return createExecution(app.fnName), nil
}
