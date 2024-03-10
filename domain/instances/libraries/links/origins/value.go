package origins

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/links/origins/resources"
)

type value struct {
	hash     hash.Hash
	resource resources.Resource
	origin   Origin
}

func createValueWithResource(
	hash hash.Hash,
	resource resources.Resource,
) Value {
	return createValueInternally(hash, resource, nil)
}

func createValueWithOrigin(
	hash hash.Hash,
	origin Origin,
) Value {
	return createValueInternally(hash, nil, origin)
}

func createValueInternally(
	hash hash.Hash,
	resource resources.Resource,
	origin Origin,
) Value {
	out := value{
		hash:     hash,
		resource: resource,
		origin:   origin,
	}

	return &out
}

// Hash returns the hash
func (obj *value) Hash() hash.Hash {
	return obj.hash
}

// IsResource returns true if there is a resource, false otherwise
func (obj *value) IsResource() bool {
	return obj.resource != nil
}

// Resource returns the resource, if any
func (obj *value) Resource() resources.Resource {
	return obj.resource
}

// IsOrigin returns true if there is an origin, false otherwise
func (obj *value) IsOrigin() bool {
	return obj.origin != nil
}

// Origin returns the origin, if any
func (obj *value) Origin() Origin {
	return obj.origin
}
