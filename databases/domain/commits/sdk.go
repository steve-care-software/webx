package commits

import (
	"math/big"
	"time"

	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
	"github.com/steve-care-software/webx/databases/domain/cryptography/hashtrees"
)

// NewBuilder creates a new builder instance
func NewBuilder(miningValue byte) Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter, miningValue)
}

// Builder represents a commit builder
type Builder interface {
	Create() Builder
	WithValues(values hashtrees.HashTree) Builder
	WithParent(parent Commit) Builder
	WithProof(proof *big.Int) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Commit, error)
}

// Commit represents a commit
type Commit interface {
	Hash() hash.Hash
	Height() uint
	Values() hashtrees.HashTree
	CreatedOn() time.Time
	HasMine() bool
	Mine() Mine
	HasParent() bool
	Parent() Commit
}

// Mine represents a mine
type Mine interface {
	Result() hash.Hash
	Proof() *big.Int
	Score() uint
}
