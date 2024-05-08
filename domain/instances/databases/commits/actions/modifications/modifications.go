package modifications

import "github.com/steve-care-software/datastencil/domain/hash"

type modifications struct {
	hash hash.Hash
	list []Modification
}

func createModifications(
	hash hash.Hash,
	list []Modification,
) Modifications {
	out := modifications{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *modifications) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *modifications) List() []Modification {
	return obj.list
}
