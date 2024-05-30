package databases

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/databases/actions"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/databases/commits"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/databases/databases"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/databases/deletes"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/databases/modifications"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/databases/retrieves"
)

type database struct {
	hash         hash.Hash
	db           databases.Database
	commit       commits.Commit
	action       actions.Action
	modification modifications.Modification
	delete       deletes.Delete
	retrieve     retrieves.Retrieve
}

func createDatabaseWithDatabase(
	hash hash.Hash,
	db databases.Database,
) Database {
	return createDatabaseInternally(hash, db, nil, nil, nil, nil, nil)
}

func createDatabaseWithCommit(
	hash hash.Hash,
	commit commits.Commit,
) Database {
	return createDatabaseInternally(hash, nil, commit, nil, nil, nil, nil)
}

func createDatabaseWithAction(
	hash hash.Hash,
	action actions.Action,
) Database {
	return createDatabaseInternally(hash, nil, nil, action, nil, nil, nil)
}

func createDatabaseWithModification(
	hash hash.Hash,
	modification modifications.Modification,
) Database {
	return createDatabaseInternally(hash, nil, nil, nil, modification, nil, nil)
}

func createDatabaseWithDelete(
	hash hash.Hash,
	delete deletes.Delete,
) Database {
	return createDatabaseInternally(hash, nil, nil, nil, nil, delete, nil)
}

func createDatabaseWithRetrieve(
	hash hash.Hash,
	retrieve retrieves.Retrieve,
) Database {
	return createDatabaseInternally(hash, nil, nil, nil, nil, nil, retrieve)
}

func createDatabaseInternally(
	hash hash.Hash,
	db databases.Database,
	commit commits.Commit,
	action actions.Action,
	modification modifications.Modification,
	delete deletes.Delete,
	retrieve retrieves.Retrieve,
) Database {
	out := database{
		hash:         hash,
		db:           db,
		commit:       commit,
		action:       action,
		modification: modification,
		delete:       delete,
		retrieve:     retrieve,
	}

	return &out
}

// Hash returns the hash
func (obj *database) Hash() hash.Hash {
	return obj.hash
}

// IsDatabase returns true if there is a database, false otherwise
func (obj *database) IsDatabase() bool {
	return obj.db != nil
}

// Database returns the database, if any
func (obj *database) Database() databases.Database {
	return obj.db
}

// IsCommit returns true if there is a commit, false otherwise
func (obj *database) IsCommit() bool {
	return obj.commit != nil
}

// Commit returns the commit, if any
func (obj *database) Commit() commits.Commit {
	return obj.commit
}

// IsAction returns true if there is an action, false otherwise
func (obj *database) IsAction() bool {
	return obj.action != nil
}

// Action returns the action, if any
func (obj *database) Action() actions.Action {
	return obj.action
}

// IsModification returns true if there is a modification, false otherwise
func (obj *database) IsModification() bool {
	return obj.modification != nil
}

// Modification returns the modification, if any
func (obj *database) Modification() modifications.Modification {
	return obj.modification
}

// IsDelete returns true if there is a delete, false otherwise
func (obj *database) IsDelete() bool {
	return obj.delete != nil
}

// Delete returns the delete, if any
func (obj *database) Delete() deletes.Delete {
	return obj.delete
}

// IsRetrieve returns true if there is a retrieve, false otherwise
func (obj *database) IsRetrieve() bool {
	return obj.retrieve != nil
}

// Retrieve returns the retrieve, if any
func (obj *database) Retrieve() retrieves.Retrieve {
	return obj.retrieve
}
