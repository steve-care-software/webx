package blockchains

import (
	"github.com/steve-care-software/webx/domain/cryptography/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a blockchain builder
type Builder interface {
	Create() Builder
	WithReference(reference hash.Hash) Builder
	WithHead(head hash.Hash) Builder
	Now() (Blockchain, error)
}

// Blockchain represents a blockchain
type Blockchain interface {
	Reference() hash.Hash
	Head() hash.Hash
}
