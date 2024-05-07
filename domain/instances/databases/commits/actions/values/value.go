package values

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/values/transforms"
)

type value struct {
	hash      hash.Hash
	instance  instances.Instance
	transform transforms.Transform
}

func createValueWithInstance(
	hash hash.Hash,
	instance instances.Instance,
) Value {
	return createValueInternally(hash, instance, nil)
}

func createValueWithTransform(
	hash hash.Hash,
	transform transforms.Transform,
) Value {
	return createValueInternally(hash, nil, transform)
}

func createValueInternally(
	hash hash.Hash,
	instance instances.Instance,
	transform transforms.Transform,
) Value {
	out := value{
		hash:      hash,
		instance:  instance,
		transform: transform,
	}

	return &out
}

// Hash returns the hash
func (obj *value) Hash() hash.Hash {
	return obj.hash
}

// IsInstance returns true if there is an instance, false otherwise
func (obj *value) IsInstance() bool {
	return obj.instance != nil
}

// Instance returns the instance, if any
func (obj *value) Instance() instances.Instance {
	return obj.instance
}

// IsTransform returns true if there is a transform, false otherwise
func (obj *value) IsTransform() bool {
	return obj.transform != nil
}

// Transform returns the transform, if any
func (obj *value) Transform() transforms.Transform {
	return obj.transform
}
