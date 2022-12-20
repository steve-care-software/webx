package connections

import (
	"errors"
	"net/url"

	"github.com/steve-care-software/webx/databases/domain/connections/contents"
)

type connectionBuilder struct {
	pIdentifier *uint
	name        string
	contents    contents.Contents
	peers       []*url.URL
}

func createConnectionBuilder() ConnectionBuilder {
	out := connectionBuilder{
		pIdentifier: nil,
		name:        "",
		contents:    nil,
		peers:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *connectionBuilder) Create() ConnectionBuilder {
	return createConnectionBuilder()
}

// WithIdentifier adds an identifier to the builder
func (app *connectionBuilder) WithIdentifier(identifier uint) ConnectionBuilder {
	app.pIdentifier = &identifier
	return app
}

// WithName adds a name to the builder
func (app *connectionBuilder) WithName(name string) ConnectionBuilder {
	app.name = name
	return app
}

// WithContents add contents to the builder
func (app *connectionBuilder) WithContents(contents contents.Contents) ConnectionBuilder {
	app.contents = contents
	return app
}

// WithPeers add peers to the builder
func (app *connectionBuilder) WithPeers(peers []*url.URL) ConnectionBuilder {
	app.peers = peers
	return app
}

// Now builds a new Connection instance
func (app *connectionBuilder) Now() (Connection, error) {
	if app.pIdentifier == nil {
		return nil, errors.New("the identifier is mandatory in order to build a Connection instance")
	}

	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Connection instance")
	}

	if app.peers != nil && len(app.peers) <= 0 {
		app.peers = nil
	}

	if app.contents != nil && app.peers != nil {
		return createConnectionWithContentsAndPeers(*app.pIdentifier, app.name, app.contents, app.peers), nil
	}

	if app.contents != nil {
		return createConnectionWithContents(*app.pIdentifier, app.name, app.contents), nil
	}

	if app.peers != nil {
		return createConnectionWithPeers(*app.pIdentifier, app.name, app.peers), nil
	}

	return createConnection(*app.pIdentifier, app.name), nil
}
