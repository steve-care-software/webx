package cardinalities

import (
	"errors"

	"github.com/steve-care-software/webx/domain/databases/entities"
)

type builder struct {
	entity entities.Entity
	pMin   *uint
	pMax   *uint
}

func createBuilder() Builder {
	out := builder{
		entity: nil,
		pMin:   nil,
		pMax:   nil,
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

// WithMin adds a min to the builder
func (app *builder) WithMin(min uint) Builder {
	app.pMin = &min
	return app
}

// WithMax adds a max to the builder
func (app *builder) WithMax(max uint) Builder {
	app.pMax = &max
	return app
}

// Now builds a new Cardinality instance
func (app *builder) Now() (Cardinality, error) {
	if app.entity == nil {
		return nil, errors.New("the entity is mandatory in order to build a Cardinality instance")
	}

	if app.pMin == nil {
		return nil, errors.New("the minimum is mandatory in order to build a Cardinality instance")
	}

	if app.pMax != nil {
		return createCardinalityWithMax(app.entity, *app.pMin, app.pMax), nil
	}

	return createCardinality(app.entity, *app.pMin), nil
}
