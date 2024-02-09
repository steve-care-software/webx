package layers

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/identity/domain/hash"
)

type layers struct {
	hash hash.Hash
	mp   map[string]Layer
	list []Layer
}

func createLayers(
	hash hash.Hash,
	mp map[string]Layer,
	list []Layer,
) Layers {
	out := layers{
		hash: hash,
		mp:   mp,
		list: list,
	}

	return &out
}

// Hash returns the hash
func (obj *layers) Hash() hash.Hash {
	return obj.hash
}

// Fetch fetches a layer by hash
func (obj *layers) Fetch(hash hash.Hash) (Layer, error) {
	keyname := hash.String()
	if ins, ok := obj.mp[keyname]; ok {
		return ins, nil
	}

	str := fmt.Sprintf("the requested Layer (hash: %s) does not exists", keyname)
	return nil, errors.New(str)
}

// List returns the list
func (obj *layers) List() []Layer {
	return obj.list
}
