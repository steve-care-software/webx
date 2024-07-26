package states

import (
	"github.com/steve-care-software/webx/engine/databases/domain/states/containers"
)

// Adapter represents an state adapter
type Adapter interface {
	ToBytes(ins State) ([]byte, error)
	ToInstance(data []byte) (State, error)
}

// Builder represents an state builder
type Builder interface {
	Create() Builder
	WithLength(length int64) Builder
	WithContainers(containers containers.Containers) Builder
	Now() (State, error)
}

// State represents an state
type State interface {
	Length() int64
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
