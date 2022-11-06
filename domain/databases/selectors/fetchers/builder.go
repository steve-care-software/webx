package fetchers

import (
	"errors"

	"github.com/steve-care-software/webx/domain/databases/entities"
)

type builder struct {
	entity    entities.Entity
	recursive entities.Identifier
	selector  entities.Identifier
}

func createBuilder() Builder {
	out := builder{
		entity:    nil,
		recursive: nil,
		selector:  nil,
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

// WithRecursive adds a recursive to the builder
func (app *builder) WithRecursive(recursive entities.Identifier) Builder {
	app.recursive = recursive
	return app
}

// WithSelector adds a selector to the builder
func (app *builder) WithSelector(selector entities.Identifier) Builder {
	app.selector = selector
	return app
}

// Now builds a new Fetcher instance
func (app *builder) Now() (Fetcher, error) {
	if app.entity == nil {
		return nil, errors.New("the entity is mandatory in order to build a Fetcher instance")
	}

	if app.recursive != nil {
		content := createContentWithRecursive(app.recursive)
		return createFetcher(app.entity, content), nil
	}

	if app.selector != nil {
		content := createContentWithSelector(app.selector)
		return createFetcher(app.entity, content), nil
	}

	return nil, errors.New("the Fetcher is invalid")
}
