package merges

import "github.com/steve-care-software/webx/engine/cursors/domain/hash"

type merge struct {
	hash hash.Hash
	base string
	top  string
}

func createMerge(
	hash hash.Hash,
	base string,
	top string,
) Merge {
	out := merge{
		hash: hash,
		base: base,
		top:  top,
	}

	return &out
}

// Hash returns the hash
func (obj *merge) Hash() hash.Hash {
	return obj.hash
}

// Base returns the base
func (obj *merge) Base() string {
	return obj.base
}

// Top returns the top
func (obj *merge) Top() string {
	return obj.top
}
