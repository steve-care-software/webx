package actions

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewActionBuilder creates a new action builder
func NewActionBuilder() ActionBuilder {
	hashAdapter := hash.NewAdapter()
	return createActionBuilder(
		hashAdapter,
	)
}

// Adapter represents the actions adapter
type Adapter interface {
	ToBytes(ins Actions) ([]byte, error)
	ToInstance(bytes []byte) (Actions, error)
}

// Builder represents an actions builder
type Builder interface {
	Create() Builder
	WithList(list []Action) Builder
	Now() (Actions, error)
}

// Actions represents actions
type Actions interface {
	Hash() hash.Hash
	List() []Action
}

// ActionBuilder represents an action builder
type ActionBuilder interface {
	Create() ActionBuilder
	WithPath(path []string) ActionBuilder
	WithModifications(modifications modifications.Modifications) ActionBuilder
	Now() (Action, error)
}

// Action represents an action
type Action interface {
	Hash() hash.Hash
	Path() []string
	Modifications() modifications.Modifications
}
