package pointers

import "github.com/steve-care-software/datastencil/domain/hash"

type pointer struct {
	hash       hash.Hash
	path       []string
	identifier hash.Hash
}

func createPointer(
	hash hash.Hash,
	path []string,
	identifier hash.Hash,
) Pointer {
	out := pointer{
		hash:       hash,
		path:       path,
		identifier: identifier,
	}

	return &out
}

// Hash returns the hash
func (obj *pointer) Hash() hash.Hash {
	return obj.hash
}

// Path returns the path
func (obj *pointer) Path() []string {
	return obj.path
}

// Identifier returns the identifier
func (obj *pointer) Identifier() hash.Hash {
	return obj.identifier
}
