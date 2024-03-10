package skeletons

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/skeletons/connections"
	"github.com/steve-care-software/datastencil/domain/skeletons/resources"
)

type builder struct {
	resources   resources.Resources
	connections connections.Connections
	previous    Skeleton
}

func createBuilder() Builder {
	out := builder{
		resources:   nil,
		connections: nil,
		previous:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithResources add resources to the builder
func (app *builder) WithResources(resources resources.Resources) Builder {
	app.resources = resources
	return app
}

// WithConnections add connections skeleton to the builder
func (app *builder) WithConnections(connections connections.Connections) Builder {
	app.connections = connections
	return app
}

// WithPrevious add previous skeleton to the builder
func (app *builder) WithPrevious(previous Skeleton) Builder {
	app.previous = previous
	return app
}

// Now builds a new Skeleton instance
func (app *builder) Now() (Skeleton, error) {
	if app.resources == nil {
		return nil, errors.New("the resources is mandatory in order to build a Skeleton instance")
	}

	version := uint(0)
	if app.previous != nil {
		version = app.previous.Version() + 1
	}

	if app.previous != nil && app.connections != nil {
		return createSkeletonWithConnectionsAndPrevious(
			version,
			app.resources,
			app.connections,
			app.previous,
		), nil
	}

	if app.connections != nil {
		return createSkeletonWithConnections(
			version,
			app.resources,
			app.connections,
		), nil
	}

	if app.previous != nil {
		return createSkeletonWithPrevious(
			version,
			app.resources,
			app.previous,
		), nil
	}

	return createSkeleton(
		version,
		app.resources,
	), nil
}
