package modifications

import "github.com/steve-care-software/datastencil/domain/hash"

type modification struct {
	hash   hash.Hash
	insert string
	delete string
}

func createModificationWithInsert(
	hash hash.Hash,
	insert string,
) Modification {
	return createModificationInternally(hash, insert, "")
}

func createModificationWithDelete(
	hash hash.Hash,
	delete string,
) Modification {
	return createModificationInternally(hash, "", delete)
}

func createModificationInternally(
	hash hash.Hash,
	insert string,
	delete string,
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
	return obj.insert != ""
}

// Insert returns the insert, if any
func (obj *modification) Insert() string {
	return obj.insert
}

// IsDelete returns true if there is a delete, false otherwise
func (obj *modification) IsDelete() bool {
	return obj.delete != ""
}

// Delete returns the delete, if any
func (obj *modification) Delete() string {
	return obj.delete
}
