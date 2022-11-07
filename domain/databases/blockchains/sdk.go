package blockchains

import (
	"github.com/steve-care-software/webx/domain/cryptography/hash"
	"github.com/steve-care-software/webx/domain/databases/entities"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a blockchain builder
type Builder interface {
	Create() Builder
	WithEntity(entity entities.Entity) Builder
	WithReference(reference hash.Hash) Builder
	WithHead(head entities.Identifier) Builder
	Now() (Blockchain, error)
}

// Blockchain represents a blockchain
type Blockchain interface {
	Entity() entities.Entity
	Reference() hash.Hash
	Head() entities.Identifier
}
