package commits

import (
	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
	"github.com/steve-care-software/webx/databases/domain/cryptography/hashtrees"
)

const commitMinSize = hash.Size + 8 + hashtrees.MinHashtreeSize

// NewAdapter creates a new adapter
func NewAdapter() Adapter {
	hashAdapter := hash.NewAdapter()
	hashTreeAdapter := hashtrees.NewAdapter()
	builder := NewBuilder()
	return createAdapter(hashAdapter, hashTreeAdapter, builder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents a commit adapter
type Adapter interface {
	ToContent(ins Commit) ([]byte, error)
	ToCommit(content []byte) (Commit, error)
}

// Builder represents a commit builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithValues(values hashtrees.HashTree) Builder
	WithParent(parent hash.Hash) Builder
	Now() (Commit, error)
}

// Commit represents a commit
type Commit interface {
	Hash() hash.Hash
	Values() hashtrees.HashTree
	HasParent() bool
	Parent() *hash.Hash
}
