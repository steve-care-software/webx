package tokens

import (
	"errors"

	"github.com/steve-care-software/webx/domain/databases/entities"
)

type builder struct {
	entity entities.Entity
	lines  Lines
	suites Suites
}

func createBuilder() Builder {
	out := builder{
		entity: nil,
		lines:  nil,
		suites: nil,
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

// WithLines add lines to the builder
func (app *builder) WithLines(lines Lines) Builder {
	app.lines = lines
	return app
}

// WithSuites add suites to the builder
func (app *builder) WithSuites(suites Suites) Builder {
	app.suites = suites
	return app
}

// Now builds a new Token instance
func (app *builder) Now() (Token, error) {
	if app.entity == nil {
		return nil, errors.New("the entity is mandatory in order to build a Token instance")
	}

	if app.lines == nil {
		return nil, errors.New("the lines is mandatory in order to build a Token instance")
	}

	if app.suites != nil {
		return createTokenWithSuites(app.entity, app.lines, app.suites), nil
	}

	return createToken(app.entity, app.lines), nil
}
