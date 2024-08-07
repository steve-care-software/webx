package states

import "github.com/steve-care-software/webx/engine/cursors/domain/storages/pointers"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewStateBuilder creates a new state builder
func NewStateBuilder() StateBuilder {
	return createStateBuilder()
}

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
}

// StateBuilder represents a state builder
type StateBuilder interface {
	Create() StateBuilder
	WithMessage(message string) StateBuilder
	WithPointers(pointers pointers.Pointers) StateBuilder
	IsDeleted() StateBuilder
	Now() (State, error)
}

// State represents a branch state
type State interface {
	Message() string
	IsDeleted() bool
	HasPointers() bool
	Pointers() pointers.Pointers
}
