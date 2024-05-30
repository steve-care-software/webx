package bridges

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers"
)

type bridge struct {
	hash  hash.Hash
	path  []string
	layer layers.Layer
}

func createBridge(
	hash hash.Hash,
	path []string,
	layer layers.Layer,
) Bridge {
	out := bridge{
		hash:  hash,
		path:  path,
		layer: layer,
	}

	return &out
}

// Hash returns the hash
func (obj *bridge) Hash() hash.Hash {
	return obj.hash
}

// Path returns the path
func (obj *bridge) Path() []string {
	return obj.path
}

// Layer returns the layer
func (obj *bridge) Layer() layers.Layer {
	return obj.layer
}
