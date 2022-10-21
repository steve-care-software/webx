package blockchains

import (
	"time"

	"github.com/steve-care-software/webx/domain/cryptography/hash"
	"github.com/steve-care-software/webx/domain/databases/blockchains/blocks"
	"github.com/steve-care-software/webx/domain/databases/blockchains/transactions"
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
	WithPendings(pendings transactions.Transactions) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Blockchain, error)
}

// Blockchain represents a blockchain
type Blockchain interface {
	Reference() hash.Hash
	Head() blocks.Block
	CreatedOn() time.Time
	HasPendings() bool
	Pendings() transactions.Transactions
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
