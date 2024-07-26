package headers

import (
	"github.com/steve-care-software/webx/engine/databases/domain/headers/containers"
)

// Adapter represents an header adapter
type Adapter interface {
	ToBytes(ins Header) ([]byte, error)
	ToInstance(data []byte) (Header, error)
}

// Builder represents an header builder
type Builder interface {
	Create() Builder
	WithLength(length int64) Builder
	WithContainers(containers containers.Containers) Builder
	Now() (Header, error)
}

// Header represents an header
type Header interface {
	Length() int64
	HasContainers() bool
	Containers() containers.Containers
}

// Repository represents an header reposiotry
type Repository interface {
	Retrieve() (Header, error)
}

// Service represents an header service
type Service interface {
	Save() (Header, error)
}
