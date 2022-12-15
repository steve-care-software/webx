package blockchains

import (
	"github.com/steve-care-software/webx/domain/blockchains/blocks"
	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a blockchain builder
type Builder interface {
	Create() Builder
	WithReference(reference hash.Hash) Builder
	WithHead(head blocks.Block) Builder
	Now() (Blockchain, error)
}

// Blockchain represents a blockchain
type Blockchain interface {
	Reference() hash.Hash
	Head() blocks.Block
}

// Repository represents a blockchain repository
type Repository interface {
	List() ([]hash.Hash, error)
	Retrieve(ref hash.Hash) (Blockchain, error)
}

// Service represents a blockchain service
type Service interface {
	Save(blockchain Blockchain) error
	Delete(blockchain Blockchain) error
}
