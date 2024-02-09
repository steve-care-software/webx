package links

import "github.com/steve-care-software/identity/domain/hash"

type originValue struct {
	hash     hash.Hash
	resource OriginResource
	origin   Origin
}

func createOriginValueWithResource(
	hash hash.Hash,
	resource OriginResource,
) OriginValue {
	return createOriginValueInternally(hash, resource, nil)
}

func createOriginValueWithOrigin(
	hash hash.Hash,
	origin Origin,
) OriginValue {
	return createOriginValueInternally(hash, nil, origin)
}

func createOriginValueInternally(
	hash hash.Hash,
	resource OriginResource,
	origin Origin,
) OriginValue {
	out := originValue{
		hash:     hash,
		resource: resource,
		origin:   origin,
	}

	return &out
}

// Hash returns the hash
func (obj *originValue) Hash() hash.Hash {
	return obj.hash
}

// IsResource returns true if there is a resource, false otherwise
func (obj *originValue) IsResource() bool {
	return obj.resource != nil
}

// Resource returns the resource, if any
func (obj *originValue) Resource() OriginResource {
	return obj.resource
}

// IsOrigin returns true if there is an origin, false otherwise
func (obj *originValue) IsOrigin() bool {
	return obj.origin != nil
}

// Origin returns the origin, if any
func (obj *originValue) Origin() Origin {
	return obj.origin
}
