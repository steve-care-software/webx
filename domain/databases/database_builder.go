package databases

import (
	"errors"
	"net/url"
)

type databaseBuilder struct {
	content     Content
	connections []url.URL
}

func createDatabaseBuilder() DatabaseBuilder {
	out := databaseBuilder{
		content:     nil,
		connections: nil,
	}

	return &out
}

// Create initializes the builder
func (app *databaseBuilder) Create() DatabaseBuilder {
	return createDatabaseBuilder()
}

// WithContent adds a content to the builder
func (app *databaseBuilder) WithContent(content Content) DatabaseBuilder {
	app.content = content
	return app
}

// WithConnections add connections to the builder
func (app *databaseBuilder) WithConnections(connections []url.URL) DatabaseBuilder {
	app.connections = connections
	return app
}

// Now builds a new Database instance
func (app *databaseBuilder) Now() (Database, error) {
	if app.content == nil {
		return nil, errors.New("the content is mandatory in order to build a Database instance")
	}

	if app.connections != nil && len(app.connections) <= 0 {
		app.connections = nil
	}

	if app.connections != nil {
		return createDatabaseWithConnections(app.content, app.connections), nil
	}

	return createDatabase(app.content), nil
}
