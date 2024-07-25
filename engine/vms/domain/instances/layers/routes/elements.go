package routes

import "github.com/steve-care-software/webx/engine/states/domain/hash"

type elements struct {
	hash hash.Hash
	list []Element
}

func createElements(
	hash hash.Hash,
	list []Element,
) Elements {
	out := elements{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *elements) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *elements) List() []Element {
	return obj.list
}
