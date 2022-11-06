package everythings

import (
	"errors"

	"github.com/steve-care-software/webx/domain/databases/entities"
)

type builder struct {
	entity    entities.Entity
	exception entities.Identifier
	escape    entities.Identifier
}

func createBuilder() Builder {
	out := builder{
		entity:    nil,
		exception: nil,
		escape:    nil,
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

// WithException adds an exception to the builder
func (app *builder) WithException(exception entities.Identifier) Builder {
	app.exception = exception
	return app
}

// WithEscape adds an escape to the builder
func (app *builder) WithEscape(escape entities.Identifier) Builder {
	app.escape = escape
	return app
}

// Now builds a new Everything instance
func (app *builder) Now() (Everything, error) {
	if app.entity == nil {
		return nil, errors.New("the entity is mandatory in order to build an Everything instance")
	}

	if app.exception == nil {
		return nil, errors.New("the exception is mandatory in order to build an Everything instance")
	}

	if app.escape != nil {
		return createEverythingWithEscape(app.entity, app.exception, app.escape), nil
	}

	return createEverything(app.entity, app.exception), nil
}
