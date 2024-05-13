package commits

import (
	"time"

	"github.com/steve-care-software/datastencil/domain/hash"
)

// Builder represents a commit builder
type Builder interface {
	Create() Builder
	WithDescription(description string) Builder
	WithActions(actions string) Builder
	WithParent(parent string) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Commit, error)
}

// Commit represents a commit
type Commit interface {
	Hash() hash.Hash
	Content() Content
	HasParent() bool
	Parent() string
}

// Content represents a commit content
type Content interface {
	Description() string
	Actions() string
	CreatedOn() time.Time
}
