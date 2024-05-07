package actions

import "github.com/steve-care-software/datastencil/domain/hash"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents the action builder
type Builder interface {
	Create() Builder
	WithPath(path string) Builder
	WithInsert(insert string) Builder
	IsDelete() Builder
	Now() (Action, error)
}

// Action represents an action
type Action interface {
	Hash() hash.Hash
	Path() string
	IsDelete() bool
	IsInsert() bool
	Insert() string
}
