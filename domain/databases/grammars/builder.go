package grammars

import (
	"errors"

	"github.com/steve-care-software/webx/domain/databases/entities"
)

type builder struct {
	entity   entities.Entity
	root     entities.Identifier
	channels entities.Identifiers
}

func createBuilder() Builder {
	out := builder{
		entity:   nil,
		root:     nil,
		channels: nil,
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

// WithRoot adds a root to the builder
func (app *builder) WithRoot(root entities.Identifier) Builder {
	app.root = root
	return app
}

// WithChannels add channels to the builder
func (app *builder) WithChannels(channels entities.Identifiers) Builder {
	app.channels = channels
	return app
}

// Now builds a new Grammar instance
func (app *builder) Now() (Grammar, error) {
	if app.entity == nil {
		return nil, errors.New("the entity is mandatory in order to build a Grammar instance")
	}

	if app.root == nil {
		return nil, errors.New("the root is mandatory in order to build a Grammar instance")
	}

	if app.channels != nil {
		return createGrammarWithChannels(app.entity, app.root, app.channels), nil
	}

	return createGrammar(app.entity, app.root), nil
}
