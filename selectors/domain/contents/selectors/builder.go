package selectors

import (
	"errors"

	"github.com/steve-care-software/webx/domain/databases/entities"
)

type builder struct {
	entity entities.Entity
	token  entities.Identifier
	inside entities.Identifier
	fn     entities.Identifier
}

func createBuilder() Builder {
	out := builder{
		entity: nil,
		token:  nil,
		inside: nil,
		fn:     nil,
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

// WithToken add token to the builder
func (app *builder) WithToken(token entities.Identifier) Builder {
	app.token = token
	return app
}

// WithInside add inside to the builder
func (app *builder) WithInside(inside entities.Identifier) Builder {
	app.inside = inside
	return app
}

// WithFunc add fn to the builder
func (app *builder) WithFunc(fn entities.Identifier) Builder {
	app.fn = fn
	return app
}

// Now builds a new Selector instance
func (app *builder) Now() (Selector, error) {
	if app.entity == nil {
		return nil, errors.New("the entity is mandatory in order to build a Selector instance")
	}

	if app.token == nil {
		return nil, errors.New("the token is mandatory in order to build a Selector instance")
	}

	if app.inside == nil {
		return nil, errors.New("the inside is mandatory in order to build a Selector instance")
	}

	if app.fn == nil {
		return nil, errors.New("the func is mandatory in order to build a Selector instance")
	}

	return createSelector(app.entity, app.token, app.inside, app.fn), nil
}
