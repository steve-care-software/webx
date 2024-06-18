package conditions

import (
	"github.com/steve-care-software/datastencil/domain/hash"
)

type condition struct {
	hash        hash.Hash
	resource    Resource
	comparisons Comparisons
}

func createCondition(
	hash hash.Hash,
	resource Resource,
) Condition {
	return createConditionInternally(
		hash,
		resource,
		nil,
	)
}

func createConditionWithComparisons(
	hash hash.Hash,
	resource Resource,
	comparisons Comparisons,
) Condition {
	return createConditionInternally(
		hash,
		resource,
		comparisons,
	)
}

func createConditionInternally(
	hash hash.Hash,
	resource Resource,
	comparisons Comparisons,
) Condition {
	out := condition{
		hash:        hash,
		resource:    resource,
		comparisons: comparisons,
	}

	return &out
}

// Hash returns the hash
func (obj *condition) Hash() hash.Hash {
	return obj.hash
}

// Resource returns the resource
func (obj *condition) Resource() Resource {
	return obj.resource
}

// HasComparisons returns true if there is a comparisons, false otherwise
func (obj *condition) HasComparisons() bool {
	return obj.comparisons != nil
}

// Comparisons returns the comparisons, if any
func (obj *condition) Comparisons() Comparisons {
	return obj.comparisons
}
