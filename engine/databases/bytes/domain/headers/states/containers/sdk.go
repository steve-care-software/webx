package containers

import (
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/headers/states/containers/pointers"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewContainerBuilder creates a new container builder
func NewContainerBuilder() ContainerBuilder {
	return createContainerBuilder()
}

// Adapter represents a containers adapter
type Adapter interface {
	InstancesToBytes(ins Containers) ([]byte, error)
	BytesToInstances(data []byte) (Containers, error)
	InstanceToBytes(ins Container) ([]byte, error)
	BytesToInstance(data []byte) (Container, error)
}

// Builder represents a containers builder
type Builder interface {
	Create() Builder
	WithList(list []Container) Builder
	Now() (Containers, error)
}

// Containers represents containers
type Containers interface {
	List() []Container
}

// ContainerBuilder represents a container builder
type ContainerBuilder interface {
	Create() ContainerBuilder
	WithKeyname(keyname string) ContainerBuilder
	WithPointers(pointers pointers.Pointers) ContainerBuilder
	Now() (Container, error)
}

// Container represents a container
type Container interface {
	Keyname() string
	Pointers() pointers.Pointers
}
