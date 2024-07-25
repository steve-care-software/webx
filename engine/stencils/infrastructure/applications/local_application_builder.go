package applications

import (
	"errors"

	db_applications "github.com/steve-care-software/webx/engine/states/applications"
	"github.com/steve-care-software/webx/engine/stencils/applications"
	infrastructure_files "github.com/steve-care-software/webx/engine/stencils/infrastructure/files"
	vm_applications "github.com/steve-care-software/webx/engine/vms/applications/layers"
	"github.com/steve-care-software/webx/engine/vms/domain/contexts"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/executions"
	json_contexts "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/contexts"
	json_executions "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/executions"
	"github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/layers"
	infrastructure_memories "github.com/steve-care-software/webx/engine/vms/infrastructure/memories"
)

type localApplicationBuilder struct {
	vmmAppFactory         vm_applications.Factory
	dbAppBuilder          db_applications.Builder
	basePath              []string
	contextEndPath        []string
	commitInnerPath       []string
	chunksInnerPath       []string
	sizeInBytesToChunk    uint
	splitHashInThisAmount uint
}

func createLocalApplicationBuilder(
	dbAppBuilder db_applications.Builder,
	contextEndPath []string,
	commitInnerPath []string,
	chunksInnerPath []string,
	sizeInBytesToChunk uint,
	splitHashInThisAmount uint,
) applications.LocalBuilder {
	out := localApplicationBuilder{
		dbAppBuilder:          dbAppBuilder,
		contextEndPath:        contextEndPath,
		commitInnerPath:       commitInnerPath,
		chunksInnerPath:       chunksInnerPath,
		sizeInBytesToChunk:    sizeInBytesToChunk,
		splitHashInThisAmount: splitHashInThisAmount,
		vmmAppFactory:         nil,
		basePath:              nil,
	}

	return &out
}

func (app *localApplicationBuilder) init(
	vmmAppFactory vm_applications.Factory,
) applications.LocalBuilder {
	app.vmmAppFactory = vmmAppFactory
	return app
}

// Create initializes the builder
func (app *localApplicationBuilder) Create() applications.LocalBuilder {
	ins := createLocalApplicationBuilder(
		app.dbAppBuilder,
		app.contextEndPath,
		app.commitInnerPath,
		app.chunksInnerPath,
		app.sizeInBytesToChunk,
		app.splitHashInThisAmount,
	)

	return ins.(*localApplicationBuilder).init(app.vmmAppFactory)
}

// WithBasePath adds a basePath to the builder
func (app *localApplicationBuilder) WithBasePath(basePath []string) applications.LocalBuilder {
	app.basePath = basePath
	return app
}

// Now builds a new Application instance
func (app *localApplicationBuilder) Now() (applications.Application, error) {
	if app.basePath != nil && len(app.basePath) <= 0 {
		app.basePath = nil
	}

	if app.basePath == nil {
		return nil, errors.New("the basePath is mandatory in order to build an Application instance")
	}

	dbApp, err := app.dbAppBuilder.Create().
		WithBasePath(app.basePath).
		WithChunksInnerPath(app.chunksInnerPath).
		WithCommitInnerPath(app.commitInnerPath).
		WithSizeInBytesToChunk(app.sizeInBytesToChunk).
		WithSplitChunkHashInThisAmountForDirectory(app.splitHashInThisAmount).
		IsJSON().
		Now()

	if err != nil {
		return nil, err
	}

	vmApp := app.vmmAppFactory.Create()
	layerAdapter := layers.NewAdapter()
	executionsRepository, executionsService := infrastructure_memories.NewExecutionRepositoryAndService()
	executionsAdapter := json_executions.NewAdapter()
	executionsBuilder := executions.NewBuilder()
	contextBuilder := contexts.NewBuilder()
	contextAdapter := json_contexts.NewAdapter()

	contextRepository := infrastructure_files.NewContextRepository(
		contextAdapter,
		app.basePath,
		app.contextEndPath,
	)

	contextService := infrastructure_files.NewContextService(
		contextAdapter,
		app.basePath,
		app.contextEndPath,
	)

	return createLocalApplication(
		dbApp,
		vmApp,
		layerAdapter,
		executionsRepository,
		executionsService,
		executionsAdapter,
		executionsBuilder,
		contextBuilder,
		contextRepository,
		contextService,
	), nil
}
