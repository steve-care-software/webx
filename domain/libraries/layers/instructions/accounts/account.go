package accounts

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/accounts/inserts"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/accounts/updates"
)

type account struct {
	hash   hash.Hash
	insert inserts.Insert
	update updates.Update
	delete string
}

func createAccountWithInsert(
	hash hash.Hash,
	insert inserts.Insert,
) Account {
	return createAccountInternally(hash, insert, nil, "")
}

func createAccountWithUpdate(
	hash hash.Hash,
	update updates.Update,
) Account {
	return createAccountInternally(hash, nil, update, "")
}

func createAccountWithDelete(
	hash hash.Hash,
	delete string,
) Account {
	return createAccountInternally(hash, nil, nil, delete)
}

func createAccountInternally(
	hash hash.Hash,
	insert inserts.Insert,
	update updates.Update,
	delete string,
) Account {
	out := account{
		hash:   hash,
		insert: insert,
		update: update,
		delete: delete,
	}

	return &out
}

// Hash returns the hash
func (obj *account) Hash() hash.Hash {
	return obj.hash
}

// IsInsert returns true if there is an insert, false otherwise
func (obj *account) IsInsert() bool {
	return obj.insert != nil
}

// Insert returns the insert, if any
func (obj *account) Insert() inserts.Insert {
	return obj.insert
}

// IsUpdate returns true if there is an update, false otherwise
func (obj *account) IsUpdate() bool {
	return obj.update != nil
}

// Update returns the update, if any
func (obj *account) Update() updates.Update {
	return obj.update
}

// IsDelete returns true if there is a delete, false otherwise
func (obj *account) IsDelete() bool {
	return obj.delete != ""
}

// Delete returns the delete, if any
func (obj *account) Delete() string {
	return obj.delete
}
