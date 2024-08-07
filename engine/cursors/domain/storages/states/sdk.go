package states

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/namespaces/singles/versions/singles/workspaces/singles/branches/singles/states/singles/pointers/storages"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/originals"
)

// Adapter represents a state adapter
type Adapter interface {
	InstancesToBytes(ins States) ([]byte, error)
	BytesToInstances(data []byte) (States, []byte, error)
	InstanceToBytes(ins State) ([]byte, error)
	BytesToInstance(data []byte) (State, []byte, error)
}

// Builder represents a states builder
type Builder interface {
	Create() Builder
	WithList(list []State) Builder
	Now() (States, error)
}

// States represents states
type States interface {
	List() []State
	FetchByName(name string) (State, error)
}

// StateBuilder represents a state builder
type StateBuilder interface {
	Create() StateBuilder
	WithOriginal(original originals.Original) StateBuilder
	WithPointers(pointers storages.Storages) StateBuilder
	Now() (State, error)
}

// State represents a branch state
type State interface {
	Original() originals.Original
	HasPointers() bool
	Pointers() storages.Storages
}
