package databases

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/databases/actions"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/databases/commits"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/databases/databases"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/databases/deletes"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/databases/modifications"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/databases/retrieves"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the database adapter
type Adapter interface {
	ToBytes(ins Database) ([]byte, error)
	ToInstance(bytes []byte) (Database, error)
}

// Builder represents a database builder
type Builder interface {
	Create() Builder
	WithDatabase(database databases.Database) Builder
	WithCommit(commit commits.Commit) Builder
	WithAction(action actions.Action) Builder
	WithModification(modification modifications.Modification) Builder
	WithDelete(delete deletes.Delete) Builder
	WithRetrieve(retrieve retrieves.Retrieve) Builder
	Now() (Database, error)
}

// Database represents a database assignable
type Database interface {
	Hash() hash.Hash
	IsDatabase() bool
	Database() databases.Database
	IsCommit() bool
	Commit() commits.Commit
	IsAction() bool
	Action() actions.Action
	IsModification() bool
	Modification() modifications.Modification
	IsDelete() bool
	Delete() deletes.Delete
	IsRetrieve() bool
	Retrieve() retrieves.Retrieve
}
