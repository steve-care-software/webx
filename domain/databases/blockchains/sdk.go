package blockchains

import (
	"time"

	"github.com/steve-care-software/webx/domain/cryptography/hash"
	"github.com/steve-care-software/webx/domain/databases/entities"
)

// Builder represents a blockchain builder
type Builder interface {
	Create() Builder
	WithEntity(entity entities.Entity) Builder
	WithReference(reference hash.Hash) Builder
	WithHead(head entities.Identifier) Builder
	WithPendings(pendings entities.Identifiers) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Blockchain, error)
}

// Blockchain represents a blockchain
type Blockchain interface {
	Entity() entities.Entity
	Reference() hash.Hash
	Head() entities.Identifier
	CreatedOn() time.Time
	HasPendings() bool
	Pendings() entities.Identifiers
}
