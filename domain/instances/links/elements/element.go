package elements

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/conditions"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers"
)

type element struct {
	hash      hash.Hash
	layer     layers.Layer
	condition conditions.Condition
}

func createElement(
	hash hash.Hash,
	layer layers.Layer,
) Element {
	return createElementInternally(hash, layer, nil)
}

func createElementWithCondition(
	hash hash.Hash,
	layer layers.Layer,
	condition conditions.Condition,
) Element {
	return createElementInternally(hash, layer, condition)
}

func createElementInternally(
	hash hash.Hash,
	layer layers.Layer,
	condition conditions.Condition,
) Element {
	out := element{
		hash:      hash,
		layer:     layer,
		condition: condition,
	}

	return &out
}

// Hash returns the hash
func (obj *element) Hash() hash.Hash {
	return obj.hash
}

// Layer returns the layer
func (obj *element) Layer() layers.Layer {
	return obj.layer
}

// HasCondition returns true if there is a condition, false otheriwse
func (obj *element) HasCondition() bool {
	return obj.condition != nil
}

// Condition returns the condition, if any
func (obj *element) Condition() conditions.Condition {
	return obj.condition
}
