package skeletons

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/skeletons/connections"
	"github.com/steve-care-software/datastencil/domain/instances/skeletons/resources"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Factory represents the skeleton factory
type Factory interface {
	Create() (Skeleton, error)
}

// Builder represents the skeleton builder
type Builder interface {
	Create() Builder
	WithResources(resources resources.Resources) Builder
	WithConnections(connections connections.Connections) Builder
	WithPrevious(previous Skeleton) Builder
	Now() (Skeleton, error)
}

// Skeleton represents a skeleton
type Skeleton interface {
	Hash() hash.Hash
	Version() uint
	Resources() resources.Resources
	HasConnections() bool
	Connections() connections.Connections
	HasPrevious() bool
	Previous() Skeleton
}
