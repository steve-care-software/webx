package elements

import (
	"errors"

	"github.com/steve-care-software/webx/domain/databases/entities"
)

type builder struct {
	entity   entities.Entity
	grammar  entities.Identifier
	contents entities.Identifiers
}

func createBuilder() Builder {
	out := builder{
		entity:   nil,
		grammar:  nil,
		contents: nil,
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

// WithContents add contents to the builder
func (app *builder) WithContents(contents entities.Identifiers) Builder {
	app.contents = contents
	return app
}

// Now builds a new Element instance
func (app *builder) Now() (Element, error) {
	if app.entity == nil {
		return nil, errors.New("the entity is mandatory in order to build an Element instance")
	}

	if app.grammar == nil {
		return nil, errors.New("the grammar is mandatory in order to build an Element instance")
	}

	if app.contents == nil {
		return nil, errors.New("the contents is mandatory in order to build an Element instance")
	}

	return createElement(app.entity, app.grammar, app.contents), nil
}
