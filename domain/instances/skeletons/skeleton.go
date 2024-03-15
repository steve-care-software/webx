package skeletons

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/skeletons/connections"
	"github.com/steve-care-software/datastencil/domain/instances/skeletons/resources"
)

type skeleton struct {
	hash        hash.Hash
	version     uint
	commit      []string
	resources   resources.Resources
	connections connections.Connections
	previous    Skeleton
}

func createSkeleton(
	hash hash.Hash,
	version uint,
	commit []string,
	resources resources.Resources,
) Skeleton {
	return createSkeletonInternally(hash, version, commit, resources, nil, nil)
}

func createSkeletonWithConnections(
	hash hash.Hash,
	version uint,
	commit []string,
	resources resources.Resources,
	connections connections.Connections,
) Skeleton {
	return createSkeletonInternally(hash, version, commit, resources, connections, nil)
}

func createSkeletonWithPrevious(
	hash hash.Hash,
	version uint,
	commit []string,
	resources resources.Resources,
	previous Skeleton,
) Skeleton {
	return createSkeletonInternally(hash, version, commit, resources, nil, previous)
}

func createSkeletonWithConnectionsAndPrevious(
	hash hash.Hash,
	version uint,
	commit []string,
	resources resources.Resources,
	connections connections.Connections,
	previous Skeleton,
) Skeleton {
	return createSkeletonInternally(hash, version, commit, resources, connections, previous)
}

func createSkeletonInternally(
	hash hash.Hash,
	version uint,
	commit []string,
	resources resources.Resources,
	connections connections.Connections,
	previous Skeleton,
) Skeleton {
	out := skeleton{
		hash:        hash,
		version:     version,
		commit:      commit,
		resources:   resources,
		connections: connections,
		previous:    previous,
	}

	return &out
}

// Hash returns the hash
func (obj *skeleton) Hash() hash.Hash {
	return obj.hash
}

// Version returns the version
func (obj *skeleton) Version() uint {
	return obj.version
}

// Commit returns the commit
func (obj *skeleton) Commit() []string {
	return obj.commit
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
