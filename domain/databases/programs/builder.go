package programs

import (
	"errors"

	"github.com/steve-care-software/webx/domain/databases/entities"
)

type builder struct {
	entity       entities.Entity
	instructions entities.Identifiers
	outputs      []uint
}

func createBuilder() Builder {
	out := builder{
		entity:       nil,
		instructions: nil,
		outputs:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithEntity adds an entity to the builder
func (app *builder) WithEntity(entity entities.Entity) Builder {
	app.entity = entity
	return app
}

// WithInstructions add instructions to the builder
func (app *builder) WithInstructions(instructions entities.Identifiers) Builder {
	app.instructions = instructions
	return app
}

// WithOutputs add outputs to the builder
func (app *builder) WithOutputs(outputs []uint) Builder {
	app.outputs = outputs
	return app
}

// Now builds a new Program instance
func (app *builder) Now() (Program, error) {
	if app.entity == nil {
		return nil, errors.New("the entity is mandatory in order to build a Program instance")
	}

	if app.instructions == nil {
		return nil, errors.New("the instructions is mandatory in order to build a Program instance")
	}

	if app.outputs != nil && len(app.outputs) <= 0 {
		app.outputs = nil
	}

	if app.outputs != nil {
		return createProgramWithOutputs(app.entity, app.instructions, app.outputs), nil
	}

	return createProgram(app.entity, app.instructions), nil
}
