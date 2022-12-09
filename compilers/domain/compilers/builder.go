package compilers

import "errors"

type builder struct {
	outputs  []string
	elements Elements
}

func createBuilder() Builder {
	out := builder{
		outputs:  nil,
		elements: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithOutputs adds an outputs to the builder
func (app *builder) WithOutputs(outputs []string) Builder {
	app.outputs = outputs
	return app
}

// WithElements add elements to the builder
func (app *builder) WithElements(elements Elements) Builder {
	app.elements = elements
	return app
}

// Now builds a new Compiler instance
func (app *builder) Now() (Compiler, error) {
	if app.outputs != nil && len(app.outputs) <= 0 {
		app.outputs = nil
	}

	if app.elements == nil {
		return nil, errors.New("the elements is mandatory in order to build a Compiler instance")
	}

	if app.outputs != nil {
		return createCompilerWithOutputs(app.elements, app.outputs), nil
	}

	return createCompiler(app.elements), nil
}
