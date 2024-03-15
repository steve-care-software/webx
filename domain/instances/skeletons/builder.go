package skeletons

import (
	"errors"
	"strconv"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/skeletons/connections"
	"github.com/steve-care-software/datastencil/domain/instances/skeletons/resources"
)

type builder struct {
	hashAdapter hash.Adapter
	commit      []string
	resources   resources.Resources
	connections connections.Connections
	previous    Skeleton
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		commit:      nil,
		resources:   nil,
		connections: nil,
		previous:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithCommit add commit resource to the builder
func (app *builder) WithCommit(commit []string) Builder {
	app.commit = commit
	return app
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

	if app.commit != nil && len(app.commit) <= 0 {
		app.commit = nil
	}

	if app.commit == nil {
		return nil, errors.New("the commit is mandatory in order to build a Skeleton instance")
	}

	version := uint(0)
	if app.previous != nil {
		version = app.previous.Version() + 1
	}

	data := [][]byte{
		app.resources.Hash().Bytes(),
		[]byte(strconv.Itoa(int(version))),
	}

	for _, oneElement := range app.commit {
		data = append(data, []byte(oneElement))
	}

	if app.connections != nil {
		data = append(data, app.connections.Hash().Bytes())
	}

	if app.previous != nil {
		data = append(data, app.previous.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.previous != nil && app.connections != nil {
		return createSkeletonWithConnectionsAndPrevious(
			*pHash,
			version,
			app.commit,
			app.resources,
			app.connections,
			app.previous,
		), nil
	}

	if app.connections != nil {
		return createSkeletonWithConnections(
			*pHash,
			version,
			app.commit,
			app.resources,
			app.connections,
		), nil
	}

	if app.previous != nil {
		return createSkeletonWithPrevious(
			*pHash,
			version,
			app.commit,
			app.resources,
			app.previous,
		), nil
	}

	return createSkeleton(
		*pHash,
		version,
		app.commit,
		app.resources,
	), nil
}
