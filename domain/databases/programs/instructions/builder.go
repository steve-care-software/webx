package instructions

import (
	"errors"

	"github.com/steve-care-software/webx/domain/databases/entities"
	"github.com/steve-care-software/webx/domain/databases/programs/assignments"
)

type builder struct {
	entity     entities.Entity
	assignment assignments.Assignment
	execution  entities.Identifier
}

func createBuilder() Builder {
	out := builder{
		entity:     nil,
		assignment: nil,
		execution:  nil,
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

// WithAssignment adds an assignment to the builder
func (app *builder) WithAssignment(assignment assignments.Assignment) Builder {
	app.assignment = assignment
	return app
}

// WithExecution adds an execution to the builder
func (app *builder) WithExecution(execution entities.Identifier) Builder {
	app.execution = execution
	return app
}

// Now builds a new Instruction instance
func (app *builder) Now() (Instruction, error) {
	if app.entity == nil {
		return nil, errors.New("the entity is mandatory in order to build an Instruction instance")
	}

	if app.assignment != nil {
		content := createContentWithAssignment(app.assignment)
		return createInstruction(app.entity, content), nil
	}

	if app.execution != nil {
		content := createContentWithExecution(app.execution)
		return createInstruction(app.entity, content), nil
	}

	return nil, errors.New("the Instruction is invalid")
}
