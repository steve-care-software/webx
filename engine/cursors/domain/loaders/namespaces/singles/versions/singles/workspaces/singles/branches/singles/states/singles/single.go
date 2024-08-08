package singles

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/namespaces/singles/versions/singles/workspaces/singles/branches/singles/states/singles/pointers"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/states"
)

type single struct {
	state    states.State
	pointers pointers.Pointers
}

func createSingleInternally(
	state states.State,
	pointers pointers.Pointers,
) Single {
	out := single{
		state:    state,
		pointers: pointers,
	}

	return &out
}

// State returns the state
func (obj *single) State() states.State {
	return obj.state
}

// HasPointers returns true if there is pointers, false otherwise
func (obj *single) HasPointers() bool {
	return obj.pointers != nil
}

// Pointers returns the pointers, if any
func (obj *single) Pointers() pointers.Pointers {
	return obj.pointers
}
