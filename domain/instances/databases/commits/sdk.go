package commits

import (
	"time"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the commit adapter
type Adapter interface {
	ToBytes(ins Commit) ([]byte, error)
	ToInstance(bytes []byte) (Commit, error)
}

// Builder represents a commit builder
type Builder interface {
	Create() Builder
	WithDescription(description string) Builder
	WithActions(actions actions.Actions) Builder
	WithParent(parent hash.Hash) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Commit, error)
}

// Commit represents a commit
type Commit interface {
	Hash() hash.Hash
	Content() Content
	HasParent() bool
	Parent() hash.Hash
}

// Content represents a commit content
type Content interface {
	Description() string
	Actions() actions.Actions
	CreatedOn() time.Time
}
