package logics

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/logics/locations"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers"
)

type logic struct {
	hash     hash.Hash
	layer    layers.Layer
	location locations.Location
}

func createLogic(
	hash hash.Hash,
	layer layers.Layer,
	location locations.Location,
) Logic {
	out := logic{
		hash:     hash,
		layer:    layer,
		location: location,
	}

	return &out
}

// Hash returns the hash
func (obj *logic) Hash() hash.Hash {
	return obj.hash
}

// Layer returns the layer
func (obj *logic) Layer() layers.Layer {
	return obj.layer
}

// Location returns the location
func (obj *logic) Location() locations.Location {
	return obj.location
}
