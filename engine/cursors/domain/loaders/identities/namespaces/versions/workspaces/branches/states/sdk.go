package states

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/strings"
)

// Adapter represents a state adapter
type Adapter interface {
	InstancesToBytes(ins States) ([]byte, error)
	BytesToInstances(data []byte) (States, error)
	InstanceToBytes(ins State) ([]byte, error)
	BytesToInstance(data []byte) (State, error)
}

// Builder represents the state builder
type Builder interface {
	Create() Builder
	WithList(list []State) Builder
	Now() (States, error)
}

// States represents states
type States interface {
	List() []State
}

// StateBuilder represents the state builder
type StateBuilder interface {
	Create() StateBuilder
	WithMessage(message string) StateBuilder
	WithResources(resources strings.Strings) StateBuilder
	WithLists(lists strings.Strings) StateBuilder
	Now() (State, error)
}

// State represents the state
type State interface {
	Message() string
	HasResources() bool
	Resources() strings.Strings
	HasLists() bool
	Lists() strings.Strings
}
