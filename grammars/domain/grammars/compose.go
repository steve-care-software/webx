package grammars

import "github.com/steve-care-software/webx/databases/domain/cryptography/hash"

type compose struct {
	hash hash.Hash
	list []ComposeElement
}

func createCompose(
	hash hash.Hash,
	list []ComposeElement,
) Compose {
	out := compose{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *compose) Hash() hash.Hash {
	return obj.hash
}

// Points returns the points
func (obj *compose) Points() uint {
	points := uint(0)
	for _, oneElement := range obj.list {
		points += oneElement.Points()
	}

	return points
}

// List returns the list of elements
func (obj *compose) List() []ComposeElement {
	return obj.list
}
