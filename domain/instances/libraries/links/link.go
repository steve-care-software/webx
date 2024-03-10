package links

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/links/elements"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/links/origins"
)

type link struct {
	hash     hash.Hash
	origin   origins.Origin
	elements elements.Elements
}

func createLink(
	hash hash.Hash,
	origin origins.Origin,
	elements elements.Elements,
) Link {
	out := link{
		hash:     hash,
		origin:   origin,
		elements: elements,
	}

	return &out
}

// Hash returns the hash
func (obj *link) Hash() hash.Hash {
	return obj.hash
}

// Origin returns the origin
func (obj *link) Origin() origins.Origin {
	return obj.origin
}

// Elements returns the elements
func (obj *link) Elements() elements.Elements {
	return obj.elements
}
