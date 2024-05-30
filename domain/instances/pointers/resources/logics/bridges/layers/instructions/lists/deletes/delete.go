package deletes

import "github.com/steve-care-software/datastencil/domain/hash"

type delete struct {
	hash  hash.Hash
	list  string
	index string
}

func createDelete(
	hash hash.Hash,
	list string,
	index string,
) Delete {
	out := delete{
		hash:  hash,
		list:  list,
		index: index,
	}

	return &out
}

// Hash returns the hash
func (obj *delete) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *delete) List() string {
	return obj.list
}

// Index returns the index
func (obj *delete) Index() string {
	return obj.index
}
