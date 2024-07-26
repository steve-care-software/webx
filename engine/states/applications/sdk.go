package applications

import (
	"github.com/steve-care-software/webx/engine/databases/entities/domain/hash"
	"github.com/steve-care-software/webx/engine/states/domain/databases"
	"github.com/steve-care-software/webx/engine/states/domain/databases/commits"
	"github.com/steve-care-software/webx/engine/states/domain/databases/commits/executions"
	"github.com/steve-care-software/webx/engine/states/domain/databases/commits/executions/chunks"
	"github.com/steve-care-software/webx/engine/states/domain/databases/metadatas"
	"github.com/steve-care-software/webx/engine/states/domain/files"
)

const invalidContextErrorPattern = "the context, %d, is invalid"
const noCommitForContextErrorPattern = "there is no commit for the context %d"

// NewApplication creates a new application
func NewApplication(
	repository databases.Repository,
	service databases.Service,
	commitRepository commits.Repository,
	chunkFileRepository files.Repository,
	chunkFileService files.Service,
	minSizeToChunkInBytes uint,
	splitHashInSubDirAmount uint,
) Application {
	hashAdapter := hash.NewAdapter()
	databaseBuilder := databases.NewBuilder()
	commitBuilder := commits.NewBuilder()
	executionsBuilder := executions.NewBuilder()
	executionBuilder := executions.NewExecutionBuilder()
	metaDataBuilder := metadatas.NewBuilder()
	chunkBuilder := chunks.NewBuilder()
	return createApplication(
		hashAdapter,
		repository,
		service,
		commitRepository,
		chunkFileRepository,
		chunkFileService,
		databaseBuilder,
		commitBuilder,
		executionsBuilder,
		executionBuilder,
		metaDataBuilder,
		chunkBuilder,
		minSizeToChunkInBytes,
		splitHashInSubDirAmount,
	)
}

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithBasePath(basePath []string) Builder
	WithCommitInnerPath(commitInnerPath []string) Builder
	WithChunksInnerPath(chunksInnerPath []string) Builder
	WithSizeInBytesToChunk(sizeInBytesToChunk uint) Builder
	WithSplitChunkHashInThisAmountForDirectory(splitHashInThisAmount uint) Builder
	IsJSON() Builder
	Now() (Application, error)
}

// Application represents an application
type Application interface {
	Retrieve(path []string) (databases.Database, error)
	RetrieveCommit(commitHash hash.Hash) (commits.Commit, error)
	RetrieveChunkBytes(fingerHash hash.Hash) ([]byte, error)
	Begin(path []string) (*uint, error)
	BeginWithInit(path []string, name string, description string) (*uint, error)
	Execute(context uint, bytes []byte) error
	Batch(context uint, bytes [][]byte) error
	Commit(context uint) error
	Cancel(context uint)
	Push(context uint) error
	RollbackToPrevious(context uint) error
	RollbackToState(context uint, headCommit hash.Hash) error
}
