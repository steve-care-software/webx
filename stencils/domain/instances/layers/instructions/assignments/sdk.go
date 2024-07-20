package assignments

import (
	"github.com/steve-care-software/datastencil/states/domain/hash"
	"github.com/steve-care-software/datastencil/stencils/domain/instances/layers/instructions/assignments/assignables"
)

// NewBuilder creates a new assignment builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the assignment adapter
type Adapter interface {
	ToBytes(ins Assignment) ([]byte, error)
	ToInstance(bytes []byte) (Assignment, error)
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
