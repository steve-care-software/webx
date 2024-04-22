package actions

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/commits/actions/pointers"
	"github.com/steve-care-software/datastencil/domain/instances/commits/actions/resources"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// NewActionBuilder creates a new action builder instance
func NewActionBuilder() ActionBuilder {
	hashAdapter := hash.NewAdapter()
	return createActionBuilder(hashAdapter)
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

// Actions represents an actions
type Actions interface {
	Hash() hash.Hash
	List() []Action
}

// ActionAdapter represents the action adapter
type ActionAdapter interface {
	ToBytes(ins Action) ([]byte, error)
	ToInstance(bytes []byte) (Action, error)
}

// ActionBuilder represents an action builder
type ActionBuilder interface {
	Create() ActionBuilder
	WithInsert(insert resources.Resource) ActionBuilder
	WithDelete(delete pointers.Pointer) ActionBuilder
	Now() (Action, error)
}

// Action represents an action
type Action interface {
	Hash() hash.Hash
	HasInsert() bool
	Insert() resources.Resource
	HasDelete() bool
	Delete() pointers.Pointer
}
