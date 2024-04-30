package conditions

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/queries/conditions/resources"
)

type element struct {
	hash      hash.Hash
	condition Condition
	resource  resources.Resource
}

func createElementWithCondition(
	hash hash.Hash,
	condition Condition,
) Element {
	return createElementInternally(hash, condition, nil)
}

func createElementWithResource(
	hash hash.Hash,
	resource resources.Resource,
) Element {
	return createElementInternally(hash, nil, resource)
}

func createElementInternally(
	hash hash.Hash,
	condition Condition,
	resource resources.Resource,
) Element {
	out := element{
		hash:      hash,
		condition: condition,
		resource:  resource,
	}

	return &out
}

// Hash returns the hash
func (obj *element) Hash() hash.Hash {
	return obj.hash
}

// IsCondition returns true if there is a condition, false otherwise
func (obj *element) IsCondition() bool {
	return obj.condition != nil
}

// Condition returns the condition, if any
func (obj *element) Condition() Condition {
	return obj.condition
}

// IsResource returns true if there is a resource, false otherwise
func (obj *element) IsResource() bool {
	return obj.resource != nil
}

// Resource returns the resource, if any
func (obj *element) Resource() resources.Resource {
	return obj.resource
}
