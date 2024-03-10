package actions

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/commits/actions/pointers"
	"github.com/steve-care-software/datastencil/domain/commits/actions/resources"
)

// Actions represents an actions
type Actions interface {
	Hash() hash.Hash
	LIst() []Action
}

// Action represents an action
type Action interface {
	Hash() hash.Hash
	HasInsert() bool
	Insert() resources.Resources
	HasDelete() bool
	Delete() pointers.Pointers
}
