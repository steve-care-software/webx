package insides

import (
	"errors"

	"github.com/steve-care-software/webx/domain/databases/entities"
)

type builder struct {
	entity   entities.Entity
	fn       entities.Identifier
	fetchers entities.Identifiers
}

func createBuilder() Builder {
	out := builder{
		entity:   nil,
		fn:       nil,
		fetchers: nil,
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

// WithFn adds a fn to the builder
func (app *builder) WithFn(fn entities.Identifier) Builder {
	app.fn = fn
	return app
}

// WithFetchers adds a fetchers to the builder
func (app *builder) WithFetchers(fetchers entities.Identifiers) Builder {
	app.fetchers = fetchers
	return app
}

// Now builds a new Inside instance
func (app *builder) Now() (Inside, error) {
	if app.entity == nil {
		return nil, errors.New("the entity is mndatory in order to build an Inside instance")
	}

	if app.fn != nil {
		content := createContentWithFn(app.fn)
		return createInside(app.entity, content), nil
	}

	if app.fetchers != nil {
		content := createContentWithFetchers(app.fetchers)
		return createInside(app.entity, content), nil
	}

	return nil, errors.New("the Inside is invalid")
}
