package files

import (
	"errors"

	"github.com/steve-care-software/webx/engine/databases/bytes/applications"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/deletes"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/entries"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/modifications"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/retrievals"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states/containers"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states/containers/pointers"
	"github.com/steve-care-software/webx/engine/databases/entities/domain/hash"
)

type applicationBuilder struct {
	hashAdapter         hash.Adapter
	statesAdapter       states.Adapter
	statesBuilder       states.Builder
	stateBuilder        states.StateBuilder
	containersBuilder   containers.Builder
	containerBuilder    containers.ContainerBuilder
	pointersBuilder     pointers.Builder
	pointerBuilder      pointers.PointerBuilder
	modificationBuilder modifications.Builder
	entriesBuilder      entries.Builder
	deletesBuilder      deletes.Builder
	retrievalsBuilder   retrievals.Builder
	basepath            []string
}

func createApplicationBuilder(
	hashAdapter hash.Adapter,
	statesAdapter states.Adapter,
	statesBuilder states.Builder,
	stateBuilder states.StateBuilder,
	containersBuilder containers.Builder,
	containerBuilder containers.ContainerBuilder,
	pointersBuilder pointers.Builder,
	pointerBuilder pointers.PointerBuilder,
	modificationBuilder modifications.Builder,
	entriesBuilder entries.Builder,
	deletesBuilder deletes.Builder,
	retrievalsBuilder retrievals.Builder,
) applications.Builder {
	out := applicationBuilder{
		hashAdapter:         hashAdapter,
		statesAdapter:       statesAdapter,
		statesBuilder:       statesBuilder,
		stateBuilder:        stateBuilder,
		containersBuilder:   containersBuilder,
		containerBuilder:    containerBuilder,
		pointersBuilder:     pointersBuilder,
		pointerBuilder:      pointerBuilder,
		modificationBuilder: modificationBuilder,
		entriesBuilder:      entriesBuilder,
		deletesBuilder:      deletesBuilder,
		retrievalsBuilder:   retrievalsBuilder,
		basepath:            nil,
	}

	return &out
}

// Create initializes the builder
func (app *applicationBuilder) Create() applications.Builder {
	return createApplicationBuilder(
		app.hashAdapter,
		app.statesAdapter,
		app.statesBuilder,
		app.stateBuilder,
		app.containersBuilder,
		app.containerBuilder,
		app.pointersBuilder,
		app.pointerBuilder,
		app.modificationBuilder,
		app.entriesBuilder,
		app.deletesBuilder,
		app.retrievalsBuilder,
	)
}

// WithBasePath adds a basePath to the builder
func (app *applicationBuilder) WithBasePath(basePath []string) applications.Builder {
	app.basepath = basePath
	return app
}

// Now builds a new Application instance
func (app *applicationBuilder) Now() (applications.Application, error) {
	if app.basepath != nil && len(app.basepath) <= 0 {
		app.basepath = nil
	}

	if app.basepath == nil {
		return nil, errors.New("the basePath is mandatory in order to build an Application instance")
	}

	return createApplication(
		app.hashAdapter,
		app.statesAdapter,
		app.statesBuilder,
		app.stateBuilder,
		app.containersBuilder,
		app.containerBuilder,
		app.pointersBuilder,
		app.pointerBuilder,
		app.modificationBuilder,
		app.entriesBuilder,
		app.deletesBuilder,
		app.retrievalsBuilder,
		app.basepath,
	), nil
}
