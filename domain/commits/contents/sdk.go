package contents

import (
	"github.com/steve-care-software/datastencil/domain/commits/contents/actions"
	"github.com/steve-care-software/datastencil/domain/commits/contents/previous"
	"github.com/steve-care-software/datastencil/domain/hash"
)

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
