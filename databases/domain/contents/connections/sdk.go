package connections

import (
	"net/url"

	"github.com/steve-care-software/webx/databases/domain/connections/contents"
)

// Adapter represents a connection adapter
type Adapter interface {
	ToContent(ins Connection) ([]byte, error)
	ToConnection(content []byte) (Connection, error)
}

// Builder represents a connection builder
type Builder interface {
	Create() Builder
	WithIdentifier(identifier uint) Builder
	WithName(name string) Builder
	WithContents(contents contents.Contents) Builder
	WithPeers(peers []*url.URL) Builder
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
