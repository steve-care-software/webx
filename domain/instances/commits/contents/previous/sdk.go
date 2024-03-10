package previous

import (
	"github.com/steve-care-software/datastencil/domain/instances/commits/contents/actions"
	"github.com/steve-care-software/datastencil/domain/hash"
)

// Builder represents a previous builder
type Builder interface {
	Create() Builder
	WithRoot(root actions.Actions) Builder
	WithPrevious(previous Previous) Builder
	Now() (Previous, error)
}

// Previous represents a previous
type Previous interface {
	Hash() hash.Hash
	IsRoot() bool
	Root() actions.Actions
	IsPrevious() bool
	Previous() Previous
}
