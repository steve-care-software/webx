package states

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/namespaces/singles/versions/singles/workspaces/singles/branches/singles/states/singles"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/states"
)

// Builder represents the state builder
type Builder interface {
	Create() Builder
	WithStates(states states.States) Builder
	WithSingle(single singles.Single) Builder
	Now() (State, error)
}

// State represents a state
type State interface {
	All() states.States
	HasSingle() bool
	Single() singles.Single
}
