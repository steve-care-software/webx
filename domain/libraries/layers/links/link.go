package links

import "github.com/steve-care-software/identity/domain/hash"

type link struct {
	hash     hash.Hash
	origin   Origin
	elements Elements
}

func createLink(
	hash hash.Hash,
	origin Origin,
	elements Elements,
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
func (obj *link) Origin() Origin {
	return obj.origin
}

// Elements returns the elements
func (obj *link) Elements() Elements {
	return obj.elements
}
