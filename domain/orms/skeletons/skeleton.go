package skeletons

import (
	"github.com/steve-care-software/datastencil/domain/orms/skeletons/connections"
	"github.com/steve-care-software/datastencil/domain/orms/skeletons/resources"
)

type skeleton struct {
	version     uint
	resources   resources.Resources
	connections connections.Connections
	previous    Skeleton
}

func createSkeleton(
	version uint,
	resources resources.Resources,
) Skeleton {
	return createSkeletonInternally(version, resources, nil, nil)
}

func createSkeletonWithConnections(
	version uint,
	resources resources.Resources,
	connections connections.Connections,
) Skeleton {
	return createSkeletonInternally(version, resources, connections, nil)
}

func createSkeletonWithPrevious(
	version uint,
	resources resources.Resources,
	previous Skeleton,
) Skeleton {
	return createSkeletonInternally(version, resources, nil, previous)
}

func createSkeletonWithConnectionsAndPrevious(
	version uint,
	resources resources.Resources,
	connections connections.Connections,
	previous Skeleton,
) Skeleton {
	return createSkeletonInternally(version, resources, connections, previous)
}

func createSkeletonInternally(
	version uint,
	resources resources.Resources,
	connections connections.Connections,
	previous Skeleton,
) Skeleton {
	out := skeleton{
		version:     version,
		resources:   resources,
		connections: connections,
		previous:    previous,
	}

	return &out
}

// Version returns the version
func (obj *skeleton) Version() uint {
	return obj.version
}

// Resources returns the resources
func (obj *skeleton) Resources() resources.Resources {
	return obj.resources
}

// HasConnections returns true if there is connections, false otherwise
func (obj *skeleton) HasConnections() bool {
	return obj.connections != nil
}

// Connections returns connections, if any
func (obj *skeleton) Connections() connections.Connections {
	return obj.connections
}

// HasPrevious returns true if there is a previous skeleton, false otherwise
func (obj *skeleton) HasPrevious() bool {
	return obj.previous != nil
}

// Previous returns the previous skeleton, if any
func (obj *skeleton) Previous() Skeleton {
	return obj.previous
}
