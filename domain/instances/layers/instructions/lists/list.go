package lists

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/lists/deletes"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/lists/inserts"
)

type list struct {
	hash   hash.Hash
	insert inserts.Insert
	delete deletes.Delete
}

func createListWithInsert(
	hash hash.Hash,
	insert inserts.Insert,
) List {
	return createListInternally(hash, insert, nil)
}

func createListWithDelete(
	hash hash.Hash,
	delete deletes.Delete,
) List {
	return createListInternally(hash, nil, delete)
}

func createListInternally(
	hash hash.Hash,
	insert inserts.Insert,
	delete deletes.Delete,
) List {
	out := list{
		hash:   hash,
		insert: insert,
		delete: delete,
	}

	return &out
}

// Hash returns the hash
func (obj *list) Hash() hash.Hash {
	return obj.hash
}

// IsInsert returns true if there is an insert, false otherwise
func (obj *list) IsInsert() bool {
	return obj.insert != nil
}

// Insert returns the insert, if any
func (obj *list) Insert() inserts.Insert {
	return obj.insert
}

// IsDelete returns true if there is a delete, false otherwise
func (obj *list) IsDelete() bool {
	return obj.delete != nil
}

// Delete returns the delete, if any
func (obj *list) Delete() deletes.Delete {
	return obj.delete
}
