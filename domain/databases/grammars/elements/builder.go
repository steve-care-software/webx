package elements

import (
	"errors"

	"github.com/steve-care-software/webx/domain/databases/entities"
)

type builder struct {
	entity      entities.Entity
	cardinality entities.Identifier
	pValue      *uint8
	external    entities.Identifier
	token       entities.Identifier
	everything  entities.Identifier
	recursive   entities.Identifier
}

func createBuilder() Builder {
	out := builder{
		entity:      nil,
		cardinality: nil,
		pValue:      nil,
		external:    nil,
		token:       nil,
		everything:  nil,
		recursive:   nil,
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

// WithCardinality adds a cardinality to the builder
func (app *builder) WithCardinality(cardinality entities.Identifier) Builder {
	app.cardinality = cardinality
	return app
}

// WithValue adds a value to the builder
func (app *builder) WithValue(value uint8) Builder {
	app.pValue = &value
	return app
}

// WithExternal adds an external to the builder
func (app *builder) WithExternal(external entities.Identifier) Builder {
	app.external = external
	return app
}

// WithToken adds a token to the builder
func (app *builder) WithToken(token entities.Identifier) Builder {
	app.token = token
	return app
}

// WithEverything adds an everything to the builder
func (app *builder) WithEverything(everything entities.Identifier) Builder {
	app.everything = everything
	return app
}

// WithRecursive adds a recursive to the builder
func (app *builder) WithRecursive(recursive entities.Identifier) Builder {
	app.recursive = recursive
	return app
}

// Now builds a new Element instance
func (app *builder) Now() (Element, error) {
	if app.entity == nil {
		return nil, errors.New("the entity is mandatory in order to build an Element instance")
	}

	if app.cardinality == nil {
		return nil, errors.New("the cardinality is mandatory in order to build an Element instance")
	}

	if app.pValue != nil {
		content := createContentWithValue(app.pValue)
		return createElement(app.entity, app.cardinality, content), nil
	}

	if app.external != nil {
		content := createContentWithExternal(app.external)
		return createElement(app.entity, app.cardinality, content), nil
	}

	if app.token != nil {
		content := createContentWithToken(app.token)
		return createElement(app.entity, app.cardinality, content), nil
	}

	if app.everything != nil {
		content := createContentWithEverything(app.everything)
		return createElement(app.entity, app.cardinality, content), nil
	}

	if app.recursive != nil {
		content := createContentWithRecursive(app.recursive)
		return createElement(app.entity, app.cardinality, content), nil
	}

	return nil, errors.New("the Element is invalid")
}
