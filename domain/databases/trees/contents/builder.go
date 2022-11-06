package contents

import (
	"errors"

	"github.com/steve-care-software/webx/domain/databases/entities"
)

type builder struct {
	entity entities.Entity
	value  Value
	prefix entities.Identifiers
}

func createBuilder() Builder {
	out := builder{
		entity: nil,
		value:  nil,
		prefix: nil,
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

// WithValue adds a value to the builder
func (app *builder) WithValue(value Value) Builder {
	app.value = value
	return app
}

// WithPrefix adds a prefix to the builder
func (app *builder) WithPrefix(prefix entities.Identifiers) Builder {
	app.prefix = prefix
	return app
}

// Now builds a new Content instance
func (app *builder) Now() (Content, error) {
	if app.entity == nil {
		return nil, errors.New("the entity is mandatory in order to build a Content instance")
	}

	if app.value == nil {
		return nil, errors.New("the value is mandatory in order to build a Content instance")
	}

	if app.prefix != nil {
		return createContentWithPrefix(app.entity, app.value, app.prefix), nil
	}

	return createContent(app.entity, app.value), nil
}
