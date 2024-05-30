package references

import (
	"github.com/steve-care-software/datastencil/domain/hash"
)

type references struct {
	hash hash.Hash
	list []Reference
}

func createReferences(
	hash hash.Hash,
	list []Reference,
) References {
	out := references{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *references) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *references) List() []Reference {
	return obj.list
}
