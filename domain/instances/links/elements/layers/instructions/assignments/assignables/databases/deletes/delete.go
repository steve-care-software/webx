package deletes

import "github.com/steve-care-software/datastencil/domain/hash"

type delete struct {
	hash   hash.Hash
	index  string
	length string
}

func createDelete(
	hash hash.Hash,
	index string,
	length string,
) Delete {
	out := delete{
		hash:   hash,
		index:  index,
		length: length,
	}

	return &out
}

// Hash returns the hash
func (obj *delete) Hash() hash.Hash {
	return obj.hash
}

// Index returns the index
func (obj *delete) Index() string {
	return obj.index
}

// Length returns the length
func (obj *delete) Length() string {
	return obj.length
}
