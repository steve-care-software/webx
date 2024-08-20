package instructions

import (
	"errors"
)

type builder struct {
	list []Instruction
}

func createBuilder() Builder {
	out := builder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithList adds a list to the builder
func (app *builder) WithList(list []Instruction) Builder {
	app.list = list
	return app
}

// Now builds a new Instructions instance
func (app *builder) Now() (Instructions, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Instruction in order to build a Instructions instance")
	}

	return createInstructions(app.list), nil
}
