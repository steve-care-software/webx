package inserts

import "github.com/steve-care-software/webx/engine/databases/entities/domain/hash"

type insert struct {
	hash    hash.Hash
	list    string
	element string
}

func createInsert(
	hash hash.Hash,
	list string,
	element string,
) Insert {
	out := insert{
		hash:    hash,
		list:    list,
		element: element,
	}

	return &out
}

// Hash returns the hash
func (obj *insert) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *insert) List() string {
	return obj.list
}

// Element returns the element
func (obj *insert) Element() string {
	return obj.element
}
