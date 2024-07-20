package constants

import "github.com/steve-care-software/datastencil/states/domain/hash"

type constants struct {
	hash hash.Hash
	list []Constant
}

func createConstants(
	hash hash.Hash,
	list []Constant,
) Constants {
	out := constants{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *constants) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *constants) List() []Constant {
	return obj.list
}
