package skeletons

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/skeletons/connections"
	"github.com/steve-care-software/datastencil/domain/instances/skeletons/resources"
	"github.com/steve-care-software/datastencil/domain/instances/skeletons/scopes"
)

type skeleton struct {
	hash        hash.Hash
	version     uint
	commit      []string
	resources   resources.Resources
	blacklist   scopes.Scopes
	connections connections.Connections
	previous    Skeleton
}

func createSkeleton(
	hash hash.Hash,
	version uint,
	commit []string,
	resources resources.Resources,
) Skeleton {
	return createSkeletonInternally(hash, version, commit, resources, nil, nil, nil)
}

func createSkeletonWithConnections(
	hash hash.Hash,
	version uint,
	commit []string,
	resources resources.Resources,
	connections connections.Connections,
) Skeleton {
	return createSkeletonInternally(hash, version, commit, resources, nil, connections, nil)
}

func createSkeletonWithBlacklist(
	hash hash.Hash,
	version uint,
	commit []string,
	resources resources.Resources,
	blacklist scopes.Scopes,
) Skeleton {
	return createSkeletonInternally(hash, version, commit, resources, blacklist, nil, nil)
}

func createSkeletonWithPrevious(
	hash hash.Hash,
	version uint,
	commit []string,
	resources resources.Resources,
	previous Skeleton,
) Skeleton {
	return createSkeletonInternally(hash, version, commit, resources, nil, nil, previous)
}

func createSkeletonWithConnectionsAndPrevious(
	hash hash.Hash,
	version uint,
	commit []string,
	resources resources.Resources,
	connections connections.Connections,
	previous Skeleton,
) Skeleton {
	return createSkeletonInternally(hash, version, commit, resources, nil, connections, previous)
}

func createSkeletonWithConnectionsAndBlacklist(
	hash hash.Hash,
	version uint,
	commit []string,
	resources resources.Resources,
	connections connections.Connections,
	blacklist scopes.Scopes,
) Skeleton {
	return createSkeletonInternally(hash, version, commit, resources, blacklist, connections, nil)
}

func createSkeletonWithBlacklistAndPrevious(
	hash hash.Hash,
	version uint,
	commit []string,
	resources resources.Resources,
	blacklist scopes.Scopes,
	previous Skeleton,
) Skeleton {
	return createSkeletonInternally(hash, version, commit, resources, blacklist, nil, previous)
}

func createSkeletonWithConnectionsAndBlacklistAndPrevious(
	hash hash.Hash,
	version uint,
	commit []string,
	resources resources.Resources,
	connections connections.Connections,
	blacklist scopes.Scopes,
	previous Skeleton,
) Skeleton {
	return createSkeletonInternally(hash, version, commit, resources, blacklist, connections, previous)
}

func createSkeletonInternally(
	hash hash.Hash,
	version uint,
	commit []string,
	resources resources.Resources,
	blacklist scopes.Scopes,
	connections connections.Connections,
	previous Skeleton,
) Skeleton {
	out := skeleton{
		hash:        hash,
		version:     version,
		commit:      commit,
		resources:   resources,
		blacklist:   blacklist,
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

// HasBlacklist returns true if there is a blacklist, false otherwise
func (obj *skeleton) HasBlacklist() bool {
	return obj.blacklist != nil
}

// Blacklist returns the blacklist, if any
func (obj *skeleton) Blacklist() scopes.Scopes {
	return obj.blacklist
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
