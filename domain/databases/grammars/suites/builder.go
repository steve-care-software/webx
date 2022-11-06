package suites

import (
	"errors"

	"github.com/steve-care-software/webx/domain/databases/entities"
)

type builder struct {
	entity  entities.Entity
	isValid bool
	content []byte
}

func createBuilder() Builder {
	out := builder{
		entity:  nil,
		isValid: false,
		content: nil,
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

// WithContent adds content to the builder
func (app *builder) WithContent(content []byte) Builder {
	app.content = content
	return app
}

// IsValid returns true if valid, false otherwise
func (app *builder) IsValid() Builder {
	app.isValid = true
	return app
}

// Now builds a new Suite instance
func (app *builder) Now() (Suite, error) {
	if app.entity == nil {
		return nil, errors.New("the entity is mandatory in order to build a Suite instance")
	}

	if app.content != nil && len(app.content) <= 0 {
		app.content = nil
	}

	if app.content == nil {
		return nil, errors.New("the content is mandatory in order to build a Suite instance")
	}

	return createSuite(app.entity, app.isValid, app.content), nil
}
