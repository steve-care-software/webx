package connections

import (
	"net/url"

	"github.com/steve-care-software/webx/databases/domain/connections/contents"
)

// Builder represents a connections builder
type Builder interface {
	Create() Builder
	WithList(list []Connection) Builder
	Now() (Connections, error)
}

// Connections represents connections
type Connections interface {
	List() []Connection
	Fetch(identifier uint) (Connection, error)
}

// ConnectionBuilder represents a connection builder
type ConnectionBuilder interface {
	Create() ConnectionBuilder
	WithIdentifier(identifier uint) ConnectionBuilder
	WithName(name string) ConnectionBuilder
	WithContents(contents contents.Contents) ConnectionBuilder
	WithPeers(peers []*url.URL) ConnectionBuilder
	Now() (Connection, error)
}

// Connection represents a connection
type Connection interface {
	Identifier() uint
	Name() string
	HasContents() bool
	Contents() contents.Contents
	HasPeers() bool
	Peers() []*url.URL
}
