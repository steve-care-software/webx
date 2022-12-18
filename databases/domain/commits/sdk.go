package commits

import (
	"time"

	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
	"github.com/steve-care-software/webx/databases/domain/cryptography/hashtrees"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Builder represents a commit builder
type Builder interface {
	Create() Builder
	WithValues(values hashtrees.HashTree) Builder
	WithParent(parent Commit) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Commit, error)
}

// Commit represents a commit
type Commit interface {
	Hash() hash.Hash
	Height() uint
	Values() hashtrees.HashTree
	CreatedOn() time.Time
	HasParent() bool
	Parent() Commit
}
