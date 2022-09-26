package publics

import (
	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/syntax/domain/identity/connections"
	"github.com/steve-care-software/syntax/domain/identity/cryptography/encryptions/keys"
	"github.com/steve-care-software/syntax/domain/identity/cryptography/hash"
	"github.com/steve-care-software/syntax/domain/identity/publics/assets"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewPublicBuilder creates a new public builder
func NewPublicBuilder() PublicBuilder {
	return createPublicBuilder()
}

// Builder represents publics builder
type Builder interface {
	Create() Builder
	WithList(list []Public) Builder
	Now() (Publics, error)
}

// Publics represents public profiles
type Publics interface {
	List() []Public
}

// PublicBuilder represents a public builder
type PublicBuilder interface {
	Create() PublicBuilder
	WithID(id uuid.UUID) PublicBuilder
	WithName(name string) PublicBuilder
	WithEncryption(encryption keys.PublicKey) PublicBuilder
	WithSignature(signature hash.Hash) PublicBuilder
	WithHost(host string) PublicBuilder
	WithPort(port uint) PublicBuilder
	WithConnections(connections connections.Connections) PublicBuilder
	WithAssets(assets assets.Assets) PublicBuilder
	Now() (Public, error)
}

// Public represents a public identity
type Public interface {
	ID() uuid.UUID
	Name() string
	Encryption() keys.PublicKey
	Signature() hash.Hash
	Host() string
	Port() uint
	HasConnections() bool
	Connections() connections.Connections
	HasAssets() bool
	Assets() assets.Assets
}

// Repository represents a public repository
type Repository interface {
	RetrieveByID(id uuid.UUID) (Public, error)
	RetrieveByName(name string) (Public, error)
}

// Service represents a public service
type Service interface {
	Save(ins Public) error
	Delete(ins Public) error
}
