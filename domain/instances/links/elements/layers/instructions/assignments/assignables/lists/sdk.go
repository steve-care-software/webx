package lists

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/lists/fetches"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents a list builder
type Builder interface {
	Create() Builder
	WithFetch(fetch fetches.Fetch) Builder
	WithLength(length string) Builder
	WithCreate(create string) Builder
	Now() (List, error)
}

// List represents a list assignable
type List interface {
	Hash() hash.Hash
	IsFetch() bool
	Fetch() fetches.Fetch
	IsLength() bool
	Length() string
	IsCreate() bool
	Create() string
}
