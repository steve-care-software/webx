package actions

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/commits/contents/actions/pointers"
	"github.com/steve-care-software/datastencil/domain/instances/commits/contents/actions/resources"
)

type action struct {
	hash   hash.Hash
	insert resources.Resource
	delete pointers.Pointer
}

func createActionWithInsert(
	hash hash.Hash,
	insert resources.Resource,
) Action {
	return createActionInternally(hash, insert, nil)
}

func createActionWithDelete(
	hash hash.Hash,
	delete pointers.Pointer,
) Action {
	return createActionInternally(hash, nil, delete)
}

func createActionInternally(
	hash hash.Hash,
	insert resources.Resource,
	delete pointers.Pointer,
) Action {
	out := action{
		hash:   hash,
		insert: insert,
		delete: delete,
	}

	return &out
}

// Hash returns the hash
func (obj *action) Hash() hash.Hash {
	return obj.hash
}

// HasInsert returns true if there is an insert, false otherwise
func (obj *action) HasInsert() bool {
	return obj.insert != nil
}

// Insert returns the insert, if any
func (obj *action) Insert() resources.Resource {
	return obj.insert
}

// HasDelete returns true if there is a delete, false otherwise
func (obj *action) HasDelete() bool {
	return obj.delete != nil
}

// Delete returns the delete, if any
func (obj *action) Delete() pointers.Pointer {
	return obj.delete
}
