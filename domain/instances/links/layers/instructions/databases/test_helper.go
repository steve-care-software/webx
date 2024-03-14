package databases

import (
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/databases/deletes"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/databases/inserts"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/databases/reverts"
)

// NewDatabaseWithInsertForTests creates a new database with insert for tests
func NewDatabaseWithInsertForTests(insert inserts.Insert) Database {
	ins, err := NewBuilder().Create().WithInsert(insert).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewDatabaseWithDeleteForTests creates a new database with delete for tests
func NewDatabaseWithDeleteForTests(delete deletes.Delete) Database {
	ins, err := NewBuilder().Create().WithDelete(delete).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewDatabaseWithCommitForTests creates a new database with commit for tests
func NewDatabaseWithCommitForTests(commit string) Database {
	ins, err := NewBuilder().Create().WithCommit(commit).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewDatabaseWithCancelForTests creates a new database with cancel for tests
func NewDatabaseWithCancelForTests(cancel string) Database {
	ins, err := NewBuilder().Create().WithCancel(cancel).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewDatabaseWithRevertForTests creates a new database with revert for tests
func NewDatabaseWithRevertForTests(revert reverts.Revert) Database {
	ins, err := NewBuilder().Create().WithRevert(revert).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
