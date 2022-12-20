package connections

import (
	"net/url"

	"github.com/steve-care-software/webx/databases/domain/connections/contents"
)

type connection struct {
	identifier uint
	name       string
	contents   contents.Contents
	peers      []*url.URL
}

func createConnection(
	identifier uint,
	name string,
) Connection {
	return createConnectionInternally(identifier, name, nil, nil)
}

func createConnectionWithContents(
	identifier uint,
	name string,
	contents contents.Contents,
) Connection {
	return createConnectionInternally(identifier, name, contents, nil)
}

func createConnectionWithPeers(
	identifier uint,
	name string,
	peers []*url.URL,
) Connection {
	return createConnectionInternally(identifier, name, nil, peers)
}

func createConnectionWithContentsAndPeers(
	identifier uint,
	name string,
	contents contents.Contents,
	peers []*url.URL,
) Connection {
	return createConnectionInternally(identifier, name, contents, peers)
}

func createConnectionInternally(
	identifier uint,
	name string,
	contents contents.Contents,
	peers []*url.URL,
) Connection {
	out := connection{
		identifier: identifier,
		name:       name,
		contents:   contents,
		peers:      peers,
	}

	return &out
}

// Identifier returns the identifier
func (obj *connection) Identifier() uint {
	return obj.identifier
}

// Name returns the name
func (obj *connection) Name() string {
	return obj.name
}

// HasContents returns true if there is contents, false otherwise
func (obj *connection) HasContents() bool {
	return obj.contents != nil
}

// Contents returns the contents, if any
func (obj *connection) Contents() contents.Contents {
	return obj.contents
}

// HasPeers returns true if there is peers, false otherwise
func (obj *connection) HasPeers() bool {
	return obj.peers != nil
}

// Peers returns the peers, if any
func (obj *connection) Peers() []*url.URL {
	return obj.peers
}
