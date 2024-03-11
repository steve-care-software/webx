package databases

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/databases/deletes"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/databases/inserts"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/databases/reverts"
)

type database struct {
	hash   hash.Hash
	insert inserts.Insert
	delete deletes.Delete
	commit string
	cancel string
	revert reverts.Revert
}

func createDatabaseWithInsert(
	hash hash.Hash,
	insert inserts.Insert,
) Database {
	return createDatabaseInternally(hash, insert, nil, "", "", nil)
}

func createDatabaseWithDelete(
	hash hash.Hash,
	delete deletes.Delete,
) Database {
	return createDatabaseInternally(hash, nil, delete, "", "", nil)
}

func createDatabaseWithCommit(
	hash hash.Hash,
	commit string,
) Database {
	return createDatabaseInternally(hash, nil, nil, commit, "", nil)
}

func createDatabaseWithCancel(
	hash hash.Hash,
	cancel string,
) Database {
	return createDatabaseInternally(hash, nil, nil, "", cancel, nil)
}

func createDatabaseWithRevert(
	hash hash.Hash,
	revert reverts.Revert,
) Database {
	return createDatabaseInternally(hash, nil, nil, "", "", revert)
}

func createDatabaseInternally(
	hash hash.Hash,
	insert inserts.Insert,
	delete deletes.Delete,
	commit string,
	cancel string,
	revert reverts.Revert,
) Database {
	out := database{
		hash:   hash,
		insert: insert,
		delete: delete,
		commit: commit,
		cancel: cancel,
		revert: revert,
	}

	return &out
}

// Hash returns the hash
func (obj *database) Hash() hash.Hash {
	return obj.hash
}

// IsInsert returns true if there is an insert, false otherwise
func (obj *database) IsInsert() bool {
	return obj.insert != nil
}

// Insert returns the insert, if any
func (obj *database) Insert() inserts.Insert {
	return obj.insert
}

// IsDelete returns true if there is a delete, false otherwise
func (obj *database) IsDelete() bool {
	return obj.delete != nil
}

// Delete returns the delete, if any
func (obj *database) Delete() deletes.Delete {
	return obj.delete
}

// IsCommit returns true if there is a commit, false otherwise
func (obj *database) IsCommit() bool {
	return obj.commit != ""
}

// Commit returns the commit, if any
func (obj *database) Commit() string {
	return obj.commit
}

// IsCancel returns true if there is a cancel, false otherwise
func (obj *database) IsCancel() bool {
	return obj.cancel != ""
}

// Cancel returns the cancel, if any
func (obj *database) Cancel() string {
	return obj.cancel
}

// IsRevert returns true if there is a revert, false otherwise
func (obj *database) IsRevert() bool {
	return obj.revert != nil
}

// Revert returns the revert, if any
func (obj *database) Revert() reverts.Revert {
	return obj.revert
}
