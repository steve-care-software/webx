package links

import "github.com/steve-care-software/datastencil/domain/hash"

type element struct {
	hash      hash.Hash
	layer     hash.Hash
	condition Condition
}

func createElement(
	hash hash.Hash,
	layer hash.Hash,
) Element {
	return createElementInternally(hash, layer, nil)
}

func createElementWithCondition(
	hash hash.Hash,
	layer hash.Hash,
	condition Condition,
) Element {
	return createElementInternally(hash, layer, condition)
}

func createElementInternally(
	hash hash.Hash,
	layer hash.Hash,
	condition Condition,
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

//  Layer returns the layer
func (obj *element) Layer() hash.Hash {
	return obj.layer
}

// HasCondition returns true if there is a condition, false otheriwse
func (obj *element) HasCondition() bool {
	return obj.condition != nil
}

// Condition returns the condition, if any
func (obj *element) Condition() Condition {
	return obj.condition
}
