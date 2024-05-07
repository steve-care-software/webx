package actions

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances"
)

type content struct {
	hash     hash.Hash
	isDelete bool
	insert   instances.Instance
}

func createContentWithDelete(
	hash hash.Hash,
) Content {
	return createContentInternally(hash, true, nil)
}

func createContentWithInsert(
	hash hash.Hash,
	insert instances.Instance,
) Content {
	return createContentInternally(hash, false, insert)
}

func createContentInternally(
	hash hash.Hash,
	isDelete bool,
	insert instances.Instance,
) Content {
	out := content{
		hash:     hash,
		isDelete: isDelete,
		insert:   insert,
	}

	return &out
}

// Hash returns the hash
func (obj *content) Hash() hash.Hash {
	return obj.hash
}

// IsDelete returns true if delete, false otherwise
func (obj *content) IsDelete() bool {
	return obj.isDelete
}

// IsInsert returns true if insert, false otherwise
func (obj *content) IsInsert() bool {
	return obj.insert != nil
}

// Insert returns the insert, if any
func (obj *content) Insert() instances.Instance {
	return obj.insert
}
