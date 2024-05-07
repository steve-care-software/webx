package commits

import (
	"time"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/commits/actions"
)

// Builder represents a commit builder
type Builder interface {
	Create() Builder
	WithDescription(description string) Builder
	WithActions(actions actions.Actions) Builder
	WithParent(parent Commit) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Commit, error)
}

// Commit represents a commit
type Commit interface {
	Hash() hash.Hash
	Height() uint
	Content() Content
	HasParent() bool
	Parent() Commit
}

// Content represents a commit content
type Content interface {
	Hash() hash.Hash
	Description() string
	Actions() actions.Actions
	CreatedOn() time.Time
}

// Repository represents a repository
type Repository interface {
	Head() (Commit, error)
	Retrieve(hash hash.Hash) (Commit, error)
}

// Service represents a service
type Service interface {
	Insert(commit Commit) error
	Delete(hash hash.Hash) error
}
