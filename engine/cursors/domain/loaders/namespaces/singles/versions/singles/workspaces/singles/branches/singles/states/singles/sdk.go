package singles

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/namespaces/singles/versions/singles/workspaces/singles/branches/singles/states/singles/pointers"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/states"
)

// Builder represents the single builder
type Builder interface {
	Create() Builder
	WithState(state states.State) Builder
	WithPointers(pointers pointers.Pointers) Builder
	Now() (Single, error)
}

// Single represents the single
type Single interface {
	State() states.State
	HasPointers() bool
	Pointers() pointers.Pointers
}
