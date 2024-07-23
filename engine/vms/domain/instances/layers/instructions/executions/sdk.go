package executions

import (
	"github.com/steve-care-software/webx/engine/states/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/executions/merges"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewContentBuilder creates a new content builder instance
func NewContentBuilder() ContentBuilder {
	hashAdapter := hash.NewAdapter()
	return createContentBuilder(
		hashAdapter,
	)
}

// Adapter represents the execution adapter
type Adapter interface {
	ToBytes(ins Execution) ([]byte, error)
	ToInstance(bytes []byte) (Execution, error)
}

// Builder represents an execution builder
type Builder interface {
	Create() Builder
	WithExecutable(executable string) Builder
	WithContent(content Content) Builder
	Now() (Execution, error)
}

// Execution represents an execution
type Execution interface {
	Hash() hash.Hash
	Executable() string
	Content() Content
}

// ContentBuilder represents a content builder
type ContentBuilder interface {
	Create() ContentBuilder
	WithCommit(commit string) ContentBuilder
	WithRollback(rollback string) ContentBuilder
	WithCancel(cancel string) ContentBuilder
	WithMerge(merge merges.Merge) ContentBuilder
	Now() (Content, error)
}

// Content represents an execution content
type Content interface {
	Hash() hash.Hash
	IsCommit() bool
	Commit() string
	IsRollback() bool
	Rollback() string
	IsCancel() bool
	Cancel() string
	IsMerge() bool
	Merge() merges.Merge
}
