package blockchains

import (
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
)

const blockchainSize = hash.Size * 2

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	hashAdapter := hash.NewAdapter()
	builder := NewBuilder()
	return createAdapter(hashAdapter, builder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents the blockchain adapter
type Adapter interface {
	ToContent(ins Blockchain) ([]byte, error)
	ToBlockchain(content []byte) (Blockchain, error)
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
