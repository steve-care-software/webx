package connections

import (
	"net/url"

	"github.com/steve-care-software/webx/databases/domain/connections/contents"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewConnectionBuilder creates a new connection builder
func NewConnectionBuilder() ConnectionBuilder {
	return createConnectionBuilder()
}

// Adapter represents a connections adapter
type Adapter interface {
	ToContent(ins Connections) ([]byte, error)
	ToConnections(content []byte) (Connections, error)
}

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

// ConnectionAdapter represents a connection adapter
type ConnectionAdapter interface {
	ToContent(ins Connection) ([]byte, error)
	ToConnection(content []byte) (Connection, error)
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
