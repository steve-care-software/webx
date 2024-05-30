package actions

import "github.com/steve-care-software/datastencil/domain/hash"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the action adapter
type Adapter interface {
	ToBytes(ins Action) ([]byte, error)
	ToInstance(bytes []byte) (Action, error)
}

// Builder represents the action builder
type Builder interface {
	Create() Builder
	WithPath(path string) Builder
	WithModifications(modifications string) Builder
	Now() (Action, error)
}

// Action represents an action
type Action interface {
	Hash() hash.Hash
	Path() string
	Modifications() string
}
