package applications

import (
	"errors"

	"github.com/steve-care-software/datastencil/applications"
	applications_layers "github.com/steve-care-software/datastencil/applications/layers"
	instructions "github.com/steve-care-software/datastencil/applications/layers/instructions"
	assignments "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments"
	assignables "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables"
	assignables_bytes "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/bytes"
	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/compilers"
	assignables_constants "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/constants"
	cryptography "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/cryptography"
	cryptography_decrypts "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/cryptography/decrypts"
	cryptography_encrypts "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/cryptography/encrypts"
	cryptography_keys "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/cryptography/keys"
	cryptography_keys_encryptions "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/cryptography/keys/encryptions"
	cryptography_keys_encryptions_decrypts "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/cryptography/keys/encryptions/decrypts"
	cryptography_keys_encryptions_encrypts "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/cryptography/keys/encryptions/encrypts"
	cryptography_keys_signatures "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/cryptography/keys/signatures"
	cryptography_keys_signatures_signs "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs"
	cryptography_keys_signatures_signs_creates "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/creates"
	cryptography_keys_signatures_signs_validates "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/validates"
	cryptography_keys_signatures_votes "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes"
	cryptography_keys_signatures_votes_creates "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/creates"
	cryptography_keys_signatures_votes_validates "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/validates"
	executables "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/excutables"
	assignables_executions "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/executions"
	assignables_executions_executes "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/executions/executes"
	assignables_executions_inits "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/executions/inits"
	assignables_executions_retrieves "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/executions/retrieves"
	assignables_lists "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/lists"
	assignables_lists_fetches "github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/lists/fetches"
	instructions_executions "github.com/steve-care-software/datastencil/applications/layers/instructions/executions"
	instructions_executions_merges "github.com/steve-care-software/datastencil/applications/layers/instructions/executions/merges"
	instructions_lists "github.com/steve-care-software/datastencil/applications/layers/instructions/lists"
	instructions_lists_deletes "github.com/steve-care-software/datastencil/applications/layers/instructions/lists/deletes"
	instructions_lists_inserts "github.com/steve-care-software/datastencil/applications/layers/instructions/lists/inserts"
	"github.com/steve-care-software/datastencil/domain/contexts"
	"github.com/steve-care-software/datastencil/domain/instances/executions"
	"github.com/steve-care-software/datastencil/infrastructure/edwards25519"
	infrastructure_files "github.com/steve-care-software/datastencil/infrastructure/files"
	"github.com/steve-care-software/datastencil/infrastructure/jsons/instances"
	json_executions "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/executions"
	"github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers"
	infrastructure_memories "github.com/steve-care-software/datastencil/infrastructure/memories"
	db_applications "github.com/steve-care-software/historydb/applications"
)

type localApplicationBuilder struct {
	dbAppBuilder          db_applications.Builder
	basePath              []string
	commitInnerPath       []string
	chunksInnerPath       []string
	sizeInBytesToChunk    uint
	splitHashInThisAmount uint
}

func createLocalApplicationBuilder(
	dbAppBuilder db_applications.Builder,
	commitInnerPath []string,
	chunksInnerPath []string,
	sizeInBytesToChunk uint,
	splitHashInThisAmount uint,
) applications.LocalBuilder {
	out := localApplicationBuilder{
		dbAppBuilder:          dbAppBuilder,
		commitInnerPath:       commitInnerPath,
		chunksInnerPath:       chunksInnerPath,
		sizeInBytesToChunk:    sizeInBytesToChunk,
		splitHashInThisAmount: splitHashInThisAmount,
		basePath:              nil,
	}

	return &out
}

// Create initializes the builder
func (app *localApplicationBuilder) Create() applications.LocalBuilder {
	return createLocalApplicationBuilder(
		app.dbAppBuilder,
		app.commitInnerPath,
		app.chunksInnerPath,
		app.sizeInBytesToChunk,
		app.splitHashInThisAmount,
	)
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

	instanceAdapter := instances.NewAdapter()
	encryptor := edwards25519.NewEncryptor()
	layerApp := applications_layers.NewApplication(
		instructions.NewApplication(
			assignments.NewApplication(
				assignables.NewApplication(
					compilers.NewApplication(
						instanceAdapter,
					),
					assignables_executions.NewApplication(
						assignables_executions_executes.NewApplication(),
						assignables_executions_inits.NewApplication(),
						assignables_executions_retrieves.NewApplication(),
					),
					assignables_bytes.NewApplication(),
					assignables_constants.NewApplication(),
					cryptography.NewApplication(
						cryptography_decrypts.NewApplication(
							encryptor,
						),
						cryptography_encrypts.NewApplication(
							encryptor,
						),
						cryptography_keys.NewApplication(
							cryptography_keys_encryptions.NewApplication(
								cryptography_keys_encryptions_decrypts.NewApplication(),
								cryptography_keys_encryptions_encrypts.NewApplication(),
								keyEncryptionBitrate,
							),
							cryptography_keys_signatures.NewApplication(
								cryptography_keys_signatures_votes.NewApplication(
									cryptography_keys_signatures_votes_creates.NewApplication(),
									cryptography_keys_signatures_votes_validates.NewApplication(),
								),
								cryptography_keys_signatures_signs.NewApplication(
									cryptography_keys_signatures_signs_creates.NewApplication(),
									cryptography_keys_signatures_signs_validates.NewApplication(),
								),
							),
						),
					),
					assignables_lists.NewApplication(
						assignables_lists_fetches.NewApplication(),
					),
					executables.NewApplication(
						NewLocalApplicationBuilder(),
						NewRemoteApplicationBuilder(),
					),
				),
			),
			instructions_lists.NewApplication(
				instructions_lists_inserts.NewApplication(),
				instructions_lists_deletes.NewApplication(),
			),
			instructions_executions.NewApplication(
				instructions_executions_merges.NewApplication(),
			),
		),
		NewLayerBinaryApplication(),
	)

	layerAdapter := layers.NewAdapter()
	executionsRepository := infrastructure_memories.NewExecutionRepository()
	executionsService := infrastructure_memories.NewExecutionService()
	executionsAdapter := json_executions.NewAdapter()
	executionsBuilder := executions.NewBuilder()
	contextBuilder := contexts.NewBuilder()
	contextRepository := infrastructure_files.NewContextRepository()
	contextService := infrastructure_files.NewContextService()

	return createLocalApplication(
		dbApp,
		layerApp,
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
