package databases

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/links/layers/instructions/databases/deletes"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/links/layers/instructions/databases/inserts"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/links/layers/instructions/databases/reverts"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents a database builder
type Builder interface {
	Create() Builder
	WithInsert(insert inserts.Insert) Builder
	WithDelete(delete deletes.Delete) Builder
	WithCommit(commit string) Builder
	WithCancel(cancel string) Builder
	WithRevert(revert reverts.Revert) Builder
	Now() (Database, error)
}

// Database represents a database
type Database interface {
	Hash() hash.Hash
	IsInsert() bool
	Insert() inserts.Insert
	IsDelete() bool
	Delete() deletes.Delete
	IsCommit() bool
	Commit() string
	IsCancel() bool
	Cancel() string
	IsRevert() bool
	Revert() reverts.Revert
}
