package containers

import (
	"github.com/steve-care-software/webx/engine/containers/domain/containers/keynames"
	"github.com/steve-care-software/webx/engine/hashes/domain/hash"
)

// Adapter represents the container adapter
type Adapter interface {
	ToBytes(ins Container) ([]byte, error)
	ToInstance(data []byte) (Container, error)
}

// Builder represents a container builder
type Builder interface {
	Create() Builder
	WithKeyname(keyname string) Builder
	WithElements(elements []hash.Hash) Builder
	Now() (Container, error)
}

// Container represents a container
type Container interface {
	Hash() hash.Hash
	Keyname() keynames.Keyname
	Elements() []hash.Hash
	Subset(index uint64, length uint64) ([]hash.Hash, error)
}
