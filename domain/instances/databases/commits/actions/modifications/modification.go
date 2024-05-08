package modifications

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications/deletes"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications/updates"
)

type modification struct {
	hash   hash.Hash
	insert []byte
	update updates.Update
	delete deletes.Delete
}

func createModificationWithInsert(
	hash hash.Hash,
	insert []byte,
) Modification {
	return createModificationInternally(hash, insert, nil, nil)
}

func createModificationWithUpdate(
	hash hash.Hash,
	update updates.Update,
) Modification {
	return createModificationInternally(hash, nil, update, nil)
}

func createModificationWithDelete(
	hash hash.Hash,
	delete deletes.Delete,
) Modification {
	return createModificationInternally(hash, nil, nil, delete)
}

func createModificationInternally(
	hash hash.Hash,
	insert []byte,
	update updates.Update,
	delete deletes.Delete,
) Modification {
	out := modification{
		hash:   hash,
		insert: insert,
		update: update,
		delete: delete,
	}

	return &out
}

// Hash returns the hash
func (obj *modification) Hash() hash.Hash {
	return obj.hash
}

// IsInsert returns true if there is an insert, false otherwise
func (obj *modification) IsInsert() bool {
	return obj.insert != nil
}

// Insert returns the insert, if any
func (obj *modification) Insert() []byte {
	return obj.insert
}

// IsUpdate returns true if there is an update, false otherwise
func (obj *modification) IsUpdate() bool {
	return obj.update != nil
}

// Update returns the update, if any
func (obj *modification) Update() updates.Update {
	return obj.update
}

// IsDelete returns true if there is a delete, false otherwise
func (obj *modification) IsDelete() bool {
	return obj.delete != nil
}

// Delete returns the delete, if any
func (obj *modification) Delete() deletes.Delete {
	return obj.delete
}
