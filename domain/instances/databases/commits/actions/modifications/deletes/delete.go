package deletes

import "github.com/steve-care-software/datastencil/domain/hash"

type delete struct {
	hash   hash.Hash
	index  uint
	length uint
}

func createDelete(
	hash hash.Hash,
	index uint,
	length uint,
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
func (obj *delete) Index() uint {
	return obj.index
}

// Length returns the length
func (obj *delete) Length() uint {
	return obj.length
}
