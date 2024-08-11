package states

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/namespaces/switchers/singles/versions/switchers/singles/workspaces/switchers/singles/branches/switchers/singles/states/singles"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/namespaces/switchers/singles/versions/switchers/singles/workspaces/switchers/singles/branches/switchers/singles/states/updates"
)

// Builder represents a states builder
type Builder interface {
	Create() Builder
	WithList(list []State) Builder
	Now() (States, error)
}

// States represents states
type States interface {
	List() []State
	Messages() []string
	Fetch(index uint64) (State, error)
}

// StateBuilder represents a state builder
type StateBuilder interface {
	Create() StateBuilder
	WithOriginal(original singles.Single) StateBuilder
	WithUpdated(updated updates.Update) StateBuilder
	Now() (State, error)
}

// State represents a state switcher
type State interface {
	Current() singles.Single
	HasOriginal() bool
	Original() singles.Single
	HasUpdated() bool
	Updated() updates.Update
}
