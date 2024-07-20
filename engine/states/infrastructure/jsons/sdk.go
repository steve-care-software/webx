package jsons

import (
	"github.com/steve-care-software/webx/engine/states/domain/databases/commits"
	"github.com/steve-care-software/webx/engine/states/domain/databases/commits/executions"
	"github.com/steve-care-software/webx/engine/states/domain/databases/commits/executions/chunks"
	"github.com/steve-care-software/webx/engine/states/domain/databases/metadatas"
	"github.com/steve-care-software/webx/engine/states/domain/databases/pointers"
	"github.com/steve-care-software/webx/engine/states/domain/hash"
)

// NewCommitAdapter creates a new commit adapter
func NewCommitAdapter() commits.Adapter {
	commitBuilder := commits.NewBuilder()
	executionsBuilder := executions.NewBuilder()
	executionBuilder := executions.NewExecutionBuilder()
	chunkBuilder := chunks.NewBuilder()
	hashAdapter := hash.NewAdapter()
	return createCommitAdapter(
		commitBuilder,
		executionsBuilder,
		executionBuilder,
		chunkBuilder,
		hashAdapter,
	)
}

// NewPointerAdapter creates a new pointer adapter
func NewPointerAdapter() pointers.Adapter {
	metaDataAdapter := NewMetaDataAdapter()
	builder := pointers.NewBuilder()
	hashAdapter := hash.NewAdapter()
	return createPointerAdapter(
		metaDataAdapter.(*MetaDataAdapter),
		builder,
		hashAdapter,
	)
}

// NewMetaDataAdapter creates a new metaData adapter
func NewMetaDataAdapter() metadatas.Adapter {
	builder := metadatas.NewBuilder()
	return createMetaDataAdapter(
		builder,
	)
}
