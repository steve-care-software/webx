package lines

import (
	"errors"

	"github.com/steve-care-software/webx/domain/databases/entities"
)

type builder struct {
	entity   entities.Entity
	grammar  entities.Identifier
	elements entities.Identifiers
}

func createBuilder() Builder {
	out := builder{
		entity:   nil,
		grammar:  nil,
		elements: nil,
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

// WithElements add elements to the builder
func (app *builder) WithElements(elements entities.Identifiers) Builder {
	app.elements = elements
	return app
}

// Now builds a new Line instance
func (app *builder) Now() (Line, error) {
	if app.entity == nil {
		return nil, errors.New("the entity is mandatory in order to build a Line instance")
	}

	if app.grammar == nil {
		return nil, errors.New("the grammar is mandatory in order to build a Line instance")
	}

	if app.elements == nil {
		return nil, errors.New("the elements is mandatory in order to build a Line instance")
	}

	return createLine(app.entity, app.grammar, app.elements), nil
}
