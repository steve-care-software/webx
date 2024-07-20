package applications

import (
	"errors"

	"github.com/steve-care-software/webx/engine/states/applications"
	"github.com/steve-care-software/webx/engine/states/domain/databases"
	"github.com/steve-care-software/webx/engine/states/domain/databases/commits"
	"github.com/steve-care-software/webx/engine/states/infrastructure/files"
	"github.com/steve-care-software/webx/engine/states/infrastructure/jsons"
)

type builder struct {
	basePath              []string
	commitInnerPath       []string
	chunksInnerPath       []string
	sizeInBytesToChunk    uint
	splitHashInThisAmount uint
	isJSON                bool
}

func createBuilder() applications.Builder {
	out := builder{}
	return &out
}

// Create initializes the builder
func (app *builder) Create() applications.Builder {
	return createBuilder()
}

// WithBasePath adds a basePath to the builder
func (app *builder) WithBasePath(basePath []string) applications.Builder {
	app.basePath = basePath
	return app
}

// WithCommitInnerPath adds a commitInnerPath to the builder
func (app *builder) WithCommitInnerPath(commitInnerPath []string) applications.Builder {
	app.commitInnerPath = commitInnerPath
	return app
}

// WithChunksInnerPath adds a chunksInnerPath to the builder
func (app *builder) WithChunksInnerPath(chunksInnerPath []string) applications.Builder {
	app.chunksInnerPath = chunksInnerPath
	return app
}

// WithSizeInBytesToChunk adds a sizeInBytesToChunk to the builder
func (app *builder) WithSizeInBytesToChunk(sizeInBytesToChunk uint) applications.Builder {
	app.sizeInBytesToChunk = sizeInBytesToChunk
	return app
}

// WithSplitChunkHashInThisAmountForDirectory adds a splitChunkHashInThisAmountForDirectory to the builder
func (app *builder) WithSplitChunkHashInThisAmountForDirectory(splitHashInThisAmount uint) applications.Builder {
	app.splitHashInThisAmount = splitHashInThisAmount
	return app
}

// IsJSON flags the builder as JSON
func (app *builder) IsJSON() applications.Builder {
	app.isJSON = true
	return app
}

func (app *builder) commit(
	basePath []string,
	innerPath []string,
) (commits.Repository, commits.Service, error) {
	commitFileRepository, err := files.NewRepositoryBuilder(innerPath).Create().
		WithBasePath(basePath).
		Now()

	if err != nil {
		return nil, nil, err
	}

	commitAdapter := jsons.NewCommitAdapter()
	commitRepository := commits.NewRepository(
		commitAdapter,
		commitFileRepository,
	)

	commitFileService, err := files.NewServiceBuilder(innerPath).Create().
		WithBasePath(basePath).
		Now()

	if err != nil {
		return nil, nil, err
	}

	commitService := commits.NewService(
		commitAdapter,
		commitFileService,
	)

	return commitRepository, commitService, nil
}

func (app *builder) database(
	basePath []string,
	commitInnerPath []string,
) (databases.Repository, databases.Service, commits.Repository, commits.Service, error) {
	commitRepository, commitService, err := app.commit(basePath, commitInnerPath)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	databaseFileRepository, err := files.NewRepositoryBuilder([]string{}).Create().
		WithBasePath(basePath).
		Now()

	if err != nil {
		return nil, nil, nil, nil, err
	}

	pointerAdapter := jsons.NewPointerAdapter()
	databaseRepository := databases.NewRepository(
		databaseFileRepository,
		commitRepository,
		pointerAdapter,
	)

	databaseFileService, err := files.NewServiceBuilder([]string{}).Create().
		WithBasePath(basePath).
		Now()

	if err != nil {
		return nil, nil, nil, nil, err
	}

	databaseService := databases.NewService(
		databaseRepository,
		databaseFileService,
		commitService,
		pointerAdapter,
	)

	return databaseRepository, databaseService, commitRepository, commitService, nil
}

// Now builds a new Application instance
func (app *builder) Now() (applications.Application, error) {
	if !app.isJSON {
		return nil, errors.New("the driver (json) is mandatory in order to build an Application")
	}

	if app.basePath != nil && len(app.basePath) <= 0 {
		app.basePath = nil
	}

	if app.basePath == nil {
		return nil, errors.New("the basePath is mandatory in order to build an Application instance")
	}

	if app.commitInnerPath != nil && len(app.commitInnerPath) <= 0 {
		app.commitInnerPath = nil
	}

	if app.commitInnerPath == nil {
		return nil, errors.New("the commitInnerPath is mandatory in order to build an Application instance")
	}

	if app.chunksInnerPath != nil && len(app.chunksInnerPath) <= 0 {
		app.chunksInnerPath = nil
	}

	if app.chunksInnerPath == nil {
		return nil, errors.New("the chunksInnerPath is mandatory in order to build an Application instance")
	}

	if app.sizeInBytesToChunk <= 0 {
		return nil, errors.New("the sizeInBytesToChunk is mandatory in order to build an Application instance")
	}

	if app.splitHashInThisAmount <= 0 {
		return nil, errors.New("the splitChunkHsshInThisAmountForFirectory is mandatory in order to build an Application instance")
	}

	chunkFileRepository, err := files.NewRepositoryBuilder(app.chunksInnerPath).Create().WithBasePath(app.basePath).Now()
	if err != nil {
		return nil, err
	}

	chunkFileService, err := files.NewServiceBuilder(app.chunksInnerPath).Create().WithBasePath(app.basePath).Now()
	if err != nil {
		return nil, err
	}

	databaseRepository, databaseService, commitRepository, _, err := app.database(app.basePath, app.commitInnerPath)
	if err != nil {
		return nil, err
	}

	return applications.NewApplication(
		databaseRepository,
		databaseService,
		commitRepository,
		chunkFileRepository,
		chunkFileService,
		app.sizeInBytesToChunk,
		app.splitHashInThisAmount,
	), nil
}
