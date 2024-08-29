package frames

import "github.com/steve-care-software/webx/engine/domain/stacks/frames/variables"

type builder struct {
	variables variables.Variables
}

func createBuilder() Builder {
	out := builder{
		variables: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithVariables add variables to the builder
func (app *builder) WithVariables(variables variables.Variables) Builder {
	app.variables = variables
	return app
}

// Now builds a new Frame instance
func (app *builder) Now() (Frame, error) {
	if app.variables != nil {
		return createFrameWithVariables(app.variables), nil
	}

	return createFrame(), nil
}
