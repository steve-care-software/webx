package compilers

import "errors"

type builder struct {
	elements []Element
}

func createBuilder() Builder {
	out := builder{
		elements: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithElements add elements to the builder
func (app *builder) WithElements(elements []Element) Builder {
	app.elements = elements
	return app
}

// Now builds a new Compiler instance
func (app *builder) Now() (Compiler, error) {
	if app.elements != nil && len(app.elements) <= 0 {
		app.elements = nil
	}

	if app.elements == nil {
		return nil, errors.New("there must be at least 1 Element in order to build a Compiler instance")
	}

	return createCompiler(app.elements), nil
}
