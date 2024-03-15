package locations

import (
	"github.com/steve-care-software/datastencil/domain/hash"
)

type location struct {
	hash   hash.Hash
	single []byte
	list   [][]byte
}

func createLocationWithSingle(
	hash hash.Hash,
	single []byte,
) Location {
	return createLocationInternally(hash, single, nil)
}

func createLocationWithList(
	hash hash.Hash,
	list [][]byte,
) Location {
	return createLocationInternally(hash, nil, list)
}

func createLocationInternally(
	hash hash.Hash,
	single []byte,
	list [][]byte,
) Location {
	out := location{
		hash:   hash,
		single: single,
		list:   list,
	}

	return &out
}

// Hash returns the hash
func (obj *location) Hash() hash.Hash {
	return obj.hash
}

// IsSingle returns true if there is a single, false otherwise
func (obj *location) IsSingle() bool {
	return obj.single != nil
}

// Single returns the single, if any
func (obj *location) Single() []byte {
	return obj.single
}

// IsList returns true if there is a list, false otherwise
func (obj *location) IsList() bool {
	return obj.list != nil
}

// List returns the list, if any
func (obj *location) List() [][]byte {
	return obj.list
}
