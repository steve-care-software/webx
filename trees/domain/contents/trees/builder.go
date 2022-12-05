package trees

import (
	"errors"

	"github.com/steve-care-software/webx/domain/databases/entities"
)

type builder struct {
	entity  entities.Entity
	grammar entities.Identifier
	line    entities.Identifier
	suffix  entities.Identifiers
}

func createBuilder() Builder {
	out := builder{
		entity:  nil,
		grammar: nil,
		line:    nil,
		suffix:  nil,
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

// WithLine adds a line to the builder
func (app *builder) WithLine(line entities.Identifier) Builder {
	app.line = line
	return app
}

// WithSuffix adds a suffix to the builder
func (app *builder) WithSuffix(suffix entities.Identifiers) Builder {
	app.suffix = suffix
	return app
}

// Now builds a new Tree instance
func (app *builder) Now() (Tree, error) {
	if app.entity == nil {
		return nil, errors.New("the entity is mandatory in order to build a Tree instance")
	}

	if app.grammar == nil {
		return nil, errors.New("the grammar is mandatory in order to build a Tree instance")
	}

	if app.line == nil {
		return nil, errors.New("the line is mandatory in order to build a Tree instance")
	}

	if app.suffix != nil {
		return createTreeWithSuffix(app.entity, app.grammar, app.line, app.suffix), nil
	}

	return createTree(app.entity, app.grammar, app.line), nil
}
