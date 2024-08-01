package states

import (
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states/pointers"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states/pointers/delimiters"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewStateBuilder creates a new state builder
func NewStateBuilder() StateBuilder {
	return createStateBuilder()
}

// Adapter represents an state adapter
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

// States represents a states
type States interface {
	List() []State
	Fetch(delimiter delimiters.Delimiter) (pointers.Pointer, error)
	Subset(index uint64, length uint64) ([]pointers.Pointer, error)
}

// StateBuilder represents a state builder
type StateBuilder interface {
	Create() StateBuilder
	WithPointers(pointers pointers.Pointers) StateBuilder
	IsDeleted() StateBuilder
	Now() (State, error)
}

// State represents an state
type State interface {
	IsDeleted() bool
	HasPointers() bool
	Pointers() pointers.Pointers
}
