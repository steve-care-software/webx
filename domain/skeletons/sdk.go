package skeletons

import (
	"github.com/steve-care-software/datastencil/domain/skeletons/connections"
	"github.com/steve-care-software/datastencil/domain/skeletons/resources"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
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
	Version() uint
	Resources() resources.Resources
	HasConnections() bool
	Connections() connections.Connections
	HasPrevious() bool
	Previous() Skeleton
}
