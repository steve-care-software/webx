package routers

import (
	"errors"

	"github.com/steve-care-software/webx/domain/databases/entities"
)

type builder struct {
	entity    entities.Entity
	grammar   entities.Identifier
	selectors entities.Identifiers
	program   entities.Identifier
}

func createBuilder() Builder {
	out := builder{
		entity:    nil,
		grammar:   nil,
		selectors: nil,
		program:   nil,
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

// WithGrammar adds a grammar to the builder
func (app *builder) WithGrammar(grammar entities.Identifier) Builder {
	app.grammar = grammar
	return app
}

// WithSelectors add selectors to the builder
func (app *builder) WithSelectors(selectors entities.Identifiers) Builder {
	app.selectors = selectors
	return app
}

// WithProgram adds a program to the builder
func (app *builder) WithProgram(program entities.Identifier) Builder {
	app.program = program
	return app
}

// Now builds a new Route instance
func (app *builder) Now() (Route, error) {
	if app.entity == nil {
		return nil, errors.New("the entity is mandatory in order to build a Route instance")
	}

	if app.grammar == nil {
		return nil, errors.New("the grammar is mandatory in order to build a Route instance")
	}

	if app.selectors == nil {
		return nil, errors.New("the selectors is mandatory in order to build a Route instance")
	}

	if app.program == nil {
		return nil, errors.New("the program is mandatory in order to build a Route instance")
	}

	return createRoute(app.entity, app.grammar, app.selectors, app.program), nil
}
