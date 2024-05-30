package databases

import (
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/databases/actions"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/databases/commits"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/databases/databases"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/databases/deletes"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/databases/modifications"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/databases/retrieves"
)

// NewDatabaseWithDatabaseForTests creates a new database with database for tests
func NewDatabaseWithDatabaseForTests(db databases.Database) Database {
	ins, err := NewBuilder().WithDatabase(db).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewDatabaseWithCommitForTests creates a new database with commit for tests
func NewDatabaseWithCommitForTests(commit commits.Commit) Database {
	ins, err := NewBuilder().WithCommit(commit).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewDatabaseWithActionForTests creates a new database with action for tests
func NewDatabaseWithActionForTests(action actions.Action) Database {
	ins, err := NewBuilder().WithAction(action).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewDatabaseWithModificationForTests creates a new database with modification for tests
func NewDatabaseWithModificationForTests(modification modifications.Modification) Database {
	ins, err := NewBuilder().WithModification(modification).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewDatabaseWithDeleteForTests creates a new database with delete for tests
func NewDatabaseWithDeleteForTests(delete deletes.Delete) Database {
	ins, err := NewBuilder().WithDelete(delete).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewDatabaseWithRetrieveForTests creates a new database with retrieve for tests
func NewDatabaseWithRetrieveForTests(retrieve retrieves.Retrieve) Database {
	ins, err := NewBuilder().WithRetrieve(retrieve).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
