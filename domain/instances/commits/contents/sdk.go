package contents

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/commits/contents/actions"
	"github.com/steve-care-software/datastencil/domain/instances/commits/contents/previous"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Builder represents the content builder
type Builder interface {
	Create() Builder
	WithAction(action actions.Action) Builder
	WithPrevious(previous previous.Previous) Builder
	Now() (Content, error)
}

// Content represents a commit content
type Content interface {
	Hash() hash.Hash
	Action() actions.Action
	Previous() previous.Previous
}
