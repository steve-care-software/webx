package layers

import "github.com/steve-care-software/datastencil/domain/hash"

type layers struct {
	hash hash.Hash
	list []Layer
}

func createLayers(
	hash hash.Hash,
	list []Layer,
) Layers {
	out := layers{
		hash: hash,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *layers) Hash() hash.Hash {
	return obj.hash
}

// List returns the list
func (obj *layers) List() []Layer {
	return obj.list
}
