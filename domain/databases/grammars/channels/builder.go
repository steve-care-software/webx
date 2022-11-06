package channels

import (
	"errors"

	"github.com/steve-care-software/webx/domain/databases/entities"
)

type builder struct {
	entity entities.Entity
	token  entities.Identifier
	prev   entities.Identifier
	next   entities.Identifier
}

func createBuilder() Builder {
	out := builder{
		entity: nil,
		token:  nil,
		prev:   nil,
		next:   nil,
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

// WithToken adds a token to the builder
func (app *builder) WithToken(token entities.Identifier) Builder {
	app.token = token
	return app
}

// WithPrevious adds a previous to the builder
func (app *builder) WithPrevious(previous entities.Identifier) Builder {
	app.prev = previous
	return app
}

// WithNext adds a next to the builder
func (app *builder) WithNext(next entities.Identifier) Builder {
	app.next = next
	return app
}

// Now builds a new Channel instance
func (app *builder) Now() (Channel, error) {
	if app.entity == nil {
		return nil, errors.New("the entity is mandatory in order to build a Channel instance")
	}

	if app.token == nil {
		return nil, errors.New("the token is mandatory in order to build a Channel instance")
	}

	if app.prev != nil && app.next != nil {
		return createChannelWithPreviousAndNext(app.entity, app.token, app.prev, app.next), nil
	}

	if app.prev != nil {
		return createChannelWithPrevious(app.entity, app.token, app.prev), nil
	}

	if app.next != nil {
		return createChannelWithNext(app.entity, app.token, app.next), nil
	}

	return createChannel(app.entity, app.token), nil
}
