package connections

import (
	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/syntax/domain/identity/cryptography/encryptions/keys"
	"github.com/steve-care-software/syntax/domain/identity/publics"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewConnectionBuilder creates a new connection builder
func NewConnectionBuilder() ConnectionBuilder {
	return createConnectionBuilder()
}

// Builder represents connections builder
type Builder interface {
	Create() Builder
	WithList(list []Connection) Builder
	Now() (Connections, error)
}

// Connections represents connections
type Connections interface {
	List() []Connection
	ListExcept(id uuid.UUID) []Connection
}

// ConnectionBuilder represents a connection builder
type ConnectionBuilder interface {
	Create() ConnectionBuilder
	WithID(id uuid.UUID) ConnectionBuilder
	WithPublic(public publics.Public) ConnectionBuilder
	WithEncryption(encryption keys.PrivateKey) ConnectionBuilder
	Now() (Connection, error)
}

// Connection represents a connection
type Connection interface {
	ID() uuid.UUID
	Public() publics.Public
	Encryption() keys.PrivateKey
}
