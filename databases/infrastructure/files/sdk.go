package files

import (
	"github.com/steve-care-software/webx/databases/applications"
	"github.com/steve-care-software/webx/databases/domain/commits"
	"github.com/steve-care-software/webx/databases/domain/commits/histories"
	"github.com/steve-care-software/webx/databases/domain/connections"
	"github.com/steve-care-software/webx/databases/domain/connections/contents"
	commit_contents "github.com/steve-care-software/webx/databases/domain/contents/commits"
	"github.com/steve-care-software/webx/databases/domain/contents/references"
	"github.com/steve-care-software/webx/databases/domain/cryptography/hashtrees"
)

const fileNameExtensionDelimiter = "."
const expectedReferenceBytesLength = 8
const filePermission = 0777

// NewApplication creates a new file application instance
func NewApplication(
	miningValue byte,
	dirPath string,
	dstExtension string,
	bckExtension string,
	readChunkSize uint,
) applications.Application {
	connectionsBuilder := connections.NewBuilder()
	connectionBuilder := connections.NewConnectionBuilder()
	contentsBuilder := contents.NewBuilder()
	contentBuilder := contents.NewContentBuilder()
	commitHistoriesAdapter := histories.NewAdapter()
	commitHistoriesBuilder := histories.NewBuilder()
	commitBuilder := commits.NewBuilder(miningValue)
	commitContentAdapter := commit_contents.NewAdapter()
	commitContentBuilder := commit_contents.NewBuilder()
	referenceAdapter := references.NewAdapter()
	referenceBuilder := references.NewBuilder()
	referenceContentKeysBuilder := references.NewContentKeysBuilder()
	referenceContentKeyBuilder := references.NewContentKeyBuilder()
	referenceCommitsBuilder := references.NewCommitsBuilder()
	referenceCommitBuilder := references.NewCommitBuilder()
	referencePointerBuilder := references.NewPointerBuilder()
	hashTreeBuilder := hashtrees.NewBuilder()
	return createApplication(
		nil,
		connectionsBuilder,
		connectionBuilder,
		contentsBuilder,
		contentBuilder,
		commitHistoriesAdapter,
		commitHistoriesBuilder,
		commitBuilder,
		commitContentAdapter,
		commitContentBuilder,
		referenceAdapter,
		referenceBuilder,
		referenceContentKeysBuilder,
		referenceContentKeyBuilder,
		referenceCommitsBuilder,
		referenceCommitBuilder,
		referencePointerBuilder,
		hashTreeBuilder,
		dirPath,
		dstExtension,
		bckExtension,
		readChunkSize,
	)
}
