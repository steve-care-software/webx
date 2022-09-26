package instructions

import "errors"

type builder struct {
	instructions []Instruction
}

func createBuilder() Builder {
	out := builder{
		instructions: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithList add instructions to the builder
func (app *builder) WithList(instructions []Instruction) Builder {
	app.instructions = instructions
	return app
}

// Now builds a new Instructions instance
func (app *builder) Now() (Instructions, error) {
	if app.instructions != nil && len(app.instructions) <= 0 {
		app.instructions = nil
	}

	if app.instructions == nil {
		return nil, errors.New("there must be at least 1 Instruction in order to build an Instructions instance")
	}

	return createInstructions(app.instructions), nil
}
