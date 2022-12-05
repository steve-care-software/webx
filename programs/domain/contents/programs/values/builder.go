package values

import (
	"errors"

	"github.com/steve-care-software/webx/domain/databases/entities"
	"github.com/steve-care-software/webx/domain/databases/programs/assignments"
)

type builder struct {
	entity     entities.Entity
	pInput     *uint
	assignment assignments.Assignment
	execution  entities.Identifier
	program    entities.Identifier
}

func createBuilder() Builder {
	out := builder{
		entity:     nil,
		pInput:     nil,
		assignment: nil,
		execution:  nil,
		program:    nil,
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

// WithInput adds an input to the builder
func (app *builder) WithInput(input uint) Builder {
	app.pInput = &input
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

// WithProgram adds a program to the builder
func (app *builder) WithProgram(program entities.Identifier) Builder {
	app.program = program
	return app
}

// Now builds a new Value instance
func (app *builder) Now() (Value, error) {
	if app.entity == nil {
		return nil, errors.New("the entity is mandatory in order to build a Value instance")
	}

	if app.pInput != nil {
		content := createContentWithInput(app.pInput)
		return createValue(app.entity, content), nil
	}

	if app.assignment != nil {
		content := createContentWithAssignment(app.assignment)
		return createValue(app.entity, content), nil
	}

	if app.execution != nil {
		content := createContentWithExecution(app.execution)
		return createValue(app.entity, content), nil
	}

	if app.program != nil {
		content := createContentWithProgram(app.program)
		return createValue(app.entity, content), nil
	}

	return nil, errors.New("the Value is invalid")
}
