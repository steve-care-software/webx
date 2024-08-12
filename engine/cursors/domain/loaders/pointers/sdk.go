package pointers

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/pointers/singles"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/pointers/updates"
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
	LastActive() (State, error)
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
