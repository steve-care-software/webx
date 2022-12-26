package files

import (
	"github.com/steve-care-software/webx/databases/applications"
	"github.com/steve-care-software/webx/databases/domain/connections"
	"github.com/steve-care-software/webx/databases/domain/connections/contents"
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
	referenceAdapter := references.NewAdapter(miningValue)
	referenceBuilder := references.NewBuilder()
	referenceContentKeysBuilder := references.NewContentKeysBuilder()
	referenceContentKeyBuilder := references.NewContentKeyBuilder()
	referenceCommitsBuilder := references.NewCommitsBuilder()
	referenceCommitAdapter := references.NewCommitAdapter(miningValue)
	referenceCommitBuilder := references.NewCommitBuilder(miningValue)
	referencePointerBuilder := references.NewPointerBuilder()
	hashTreeBuilder := hashtrees.NewBuilder()
	return createApplication(
		nil,
		connectionsBuilder,
		connectionBuilder,
		contentsBuilder,
		contentBuilder,
		referenceAdapter,
		referenceBuilder,
		referenceContentKeysBuilder,
		referenceContentKeyBuilder,
		referenceCommitsBuilder,
		referenceCommitAdapter,
		referenceCommitBuilder,
		referencePointerBuilder,
		hashTreeBuilder,
		dirPath,
		dstExtension,
		bckExtension,
		readChunkSize,
	)
}
