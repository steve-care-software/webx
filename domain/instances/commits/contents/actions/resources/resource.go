package resources

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances"
)

type resource struct {
	hash     hash.Hash
	path     []string
	instance instances.Instance
}

func createResource(
	hash hash.Hash,
	path []string,
	instance instances.Instance,
) Resource {
	out := resource{
		hash:     hash,
		path:     path,
		instance: instance,
	}

	return &out
}

// Hash returns the hash
func (obj *resource) Hash() hash.Hash {
	return obj.hash
}

// Path returns the path
func (obj *resource) Path() []string {
	return obj.path
}

// Instance returns the instance
func (obj *resource) Instance() instances.Instance {
	return obj.instance
}
