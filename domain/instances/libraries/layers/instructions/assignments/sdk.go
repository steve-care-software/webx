package assignments

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables"
)

// NewBuilder creates a new assignment builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents an assignment builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithAssignable(assignable assignables.Assignable) Builder
	Now() (Assignment, error)
}

// Assignment represents an assignment
type Assignment interface {
	Hash() hash.Hash
	Name() string
	Assignable() assignables.Assignable
}
