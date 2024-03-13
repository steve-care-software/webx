package mocks

import (
	"github.com/steve-care-software/datastencil/domain/skeletons"
	"github.com/steve-care-software/datastencil/domain/skeletons/connections"
	"github.com/steve-care-software/datastencil/domain/skeletons/resources"
)

type skeleton struct {
	version uint
}

func createSkeleton() skeletons.Skeleton {
	out := skeleton{
		version: 1,
	}
	return &out
}

// Version returns the version
func (obj *skeleton) Version() uint {
	return obj.version
}

// Resources returns the resources
func (obj *skeleton) Resources() resources.Resources {
	return nil
}

// HasConnections returns true if there is connections, false otherwise
func (obj *skeleton) HasConnections() bool {
	return false
}

// Connections returns the connections, if any
func (obj *skeleton) Connections() connections.Connections {
	return nil
}

// HasPrevious returns true if there is a previous, false otherwise
func (obj *skeleton) HasPrevious() bool {
	return false
}

// Previous returns the previous, if any
func (obj *skeleton) Previous() skeletons.Skeleton {
	return nil
}
