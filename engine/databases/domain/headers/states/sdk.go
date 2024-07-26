package states

import (
	"github.com/steve-care-software/webx/engine/databases/domain/headers/states/containers"
)

// Adapter represents an state adapter
type Adapter interface {
	InstanceToBytes(ins State) ([]byte, error)
	BytesToInstance(data []byte) (State, error)
	InstancesToBytes(ins States) ([]byte, error)
	BytesToInstances(data []byte) (States, error)
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
}

// StateBuilder represents a state builder
type StateBuilder interface {
	Create() StateBuilder
	WithContainers(containers containers.Containers) StateBuilder
	IsDeleted() bool
	Now() (State, error)
}

// State represents an state
type State interface {
	IsDeleted() bool
	HasContainers() bool
	Containers() containers.Containers
}

// Repository represents an state reposiotry
type Repository interface {
	Retrieve() (State, error)
}

// Service represents an state service
type Service interface {
	Save() (State, error)
}
