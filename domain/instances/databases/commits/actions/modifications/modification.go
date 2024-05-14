package modifications

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications/deletes"
)

type modification struct {
	hash   hash.Hash
	insert []byte
	delete deletes.Delete
}

func createModificationWithInsert(
	hash hash.Hash,
	insert []byte,
) Modification {
	return createModificationInternally(hash, insert, nil)
}

func createModificationWithDelete(
	hash hash.Hash,
	delete deletes.Delete,
) Modification {
	return createModificationInternally(hash, nil, delete)
}

func createModificationInternally(
	hash hash.Hash,
	insert []byte,
	delete deletes.Delete,
) Modification {
	out := modification{
		hash:   hash,
		insert: insert,
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

// IsDelete returns true if there is a delete, false otherwise
func (obj *modification) IsDelete() bool {
	return obj.delete != nil
}

// Delete returns the delete, if any
func (obj *modification) Delete() deletes.Delete {
	return obj.delete
}