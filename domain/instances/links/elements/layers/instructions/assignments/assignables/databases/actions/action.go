package actions

import "github.com/steve-care-software/datastencil/domain/hash"

type action struct {
	hash     hash.Hash
	path     string
	insert   string
	isDelete bool
}

func createActionWithInsert(
	hash hash.Hash,
	path string,
	insert string,
) Action {
	return createActionInternally(hash, path, insert, false)
}

func createActionWithDelete(
	hash hash.Hash,
	path string,
) Action {
	return createActionInternally(hash, path, "", false)
}

func createActionInternally(
	hash hash.Hash,
	path string,
	insert string,
	isDelete bool,
) Action {
	out := action{
		hash:     hash,
		path:     path,
		insert:   insert,
		isDelete: isDelete,
	}

	return &out
}

// Hash returns the hash
func (obj *action) Hash() hash.Hash {
	return obj.hash
}

// Path returns the path
func (obj *action) Path() string {
	return obj.path
}

// IsDelete returns true if delete, false otherwise
func (obj *action) IsDelete() bool {
	return obj.isDelete
}

// IsInsert returns true if insert, false otherwise
func (obj *action) IsInsert() bool {
	return obj.insert != ""
}

// Insert returns insert, if any
func (obj *action) Insert() string {
	return obj.insert
}
