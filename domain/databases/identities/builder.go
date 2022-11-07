package identities

import (
	"errors"

	"github.com/steve-care-software/webx/domain/databases/entities"
)

type builder struct {
	entity        entities.Entity
	modifications entities.Identifiers
}

func createBuilder() Builder {
	out := builder{
		entity:        nil,
		modifications: nil,
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

// WithModifications add modifications to the builder
func (app *builder) WithModifications(modifications entities.Identifiers) Builder {
	app.modifications = modifications
	return app
}

// Now builds a new Identity instance
func (app *builder) Now() (Identity, error) {
	if app.entity == nil {
		return nil, errors.New("the entity is mandatory in order to build an Identity instance")
	}

	if app.modifications == nil {
		return nil, errors.New("the modifications is mandatory in order to build an Identity instance")
	}

	return createIdentity(app.entity, app.modifications), nil
}
