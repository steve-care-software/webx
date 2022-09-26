package publics

import (
	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/syntax/domain/identity/cryptography/encryptions/keys"
	"github.com/steve-care-software/syntax/domain/identity/cryptography/hash"
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
}
