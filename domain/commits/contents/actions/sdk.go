package actions

import (
	"github.com/steve-care-software/datastencil/domain/commits/contents/actions/pointers"
	"github.com/steve-care-software/datastencil/domain/commits/contents/actions/resources"
	"github.com/steve-care-software/datastencil/domain/hash"
)

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
