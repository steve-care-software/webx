package fns

import (
	"errors"

	"github.com/steve-care-software/webx/domain/databases/entities"
)

type builder struct {
	entity    entities.Entity
	isSingle  bool
	isContent bool
	program   entities.Identifier
	pParam    *uint
}

func createBuilder() Builder {
	out := builder{
		entity:    nil,
		isSingle:  false,
		isContent: false,
		program:   nil,
		pParam:    nil,
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

// WithProgram adds a program to the builder
func (app *builder) WithProgram(program entities.Identifier) Builder {
	app.program = program
	return app
}

// WithParam adds a param to the builder
func (app *builder) WithParam(param uint) Builder {
	app.pParam = &param
	return app
}

// IsSingle flags the builder as single
func (app *builder) IsSingle() Builder {
	app.isSingle = true
	return app
}

// IsContent flags the builder as content
func (app *builder) IsContent() Builder {
	app.isContent = true
	return app
}

// Now builds a new Fn isntance
func (app *builder) Now() (Fn, error) {
	if app.entity == nil {
		return nil, errors.New("the entity is mandatory in order to build a Fn instance")
	}

	if app.program == nil {
		return nil, errors.New("the program is mandatory in order to build a Fn instance")
	}

	if app.pParam == nil {
		return nil, errors.New("the param is mandatory in order to build a Fn instance")
	}

	return createFn(app.entity, app.isSingle, app.isContent, app.program, *app.pParam), nil
}
