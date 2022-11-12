package databases

import (
	"errors"
	"net/url"
)

type builder struct {
	head        Head
	connections []url.URL
}

func createBuilder() Builder {
	out := builder{
		head:        nil,
		connections: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithHead adds a head to the builder
func (app *builder) WithHead(head Head) Builder {
	app.head = head
	return app
}

// WithConnections add connections to the builder
func (app *builder) WithConnections(connections []url.URL) Builder {
	app.connections = connections
	return app
}

// Now builds a new Database instance
func (app *builder) Now() (Database, error) {
	if app.head == nil {
		return nil, errors.New("the head is mandatory in order to build a Database instance")
	}

	if app.connections != nil && len(app.connections) <= 0 {
		app.connections = nil
	}

	if app.connections != nil {
		return createDatabaseWithConnections(app.head, app.connections), nil
	}

	return createDatabase(app.head), nil
}
