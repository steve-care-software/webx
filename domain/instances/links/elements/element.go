package elements

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/conditions"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/logics"
)

type element struct {
	hash      hash.Hash
	logic     logics.Logic
	condition conditions.Condition
}

func createElement(
	hash hash.Hash,
	logic logics.Logic,
) Element {
	return createElementInternally(hash, logic, nil)
}

func createElementWithCondition(
	hash hash.Hash,
	logic logics.Logic,
	condition conditions.Condition,
) Element {
	return createElementInternally(hash, logic, condition)
}

func createElementInternally(
	hash hash.Hash,
	logic logics.Logic,
	condition conditions.Condition,
) Element {
	out := element{
		hash:      hash,
		logic:     logic,
		condition: condition,
	}

	return &out
}

// Hash returns the hash
func (obj *element) Hash() hash.Hash {
	return obj.hash
}

// Logic returns the logic
func (obj *element) Logic() logics.Logic {
	return obj.logic
}

// HasCondition returns true if there is a condition, false otheriwse
func (obj *element) HasCondition() bool {
	return obj.condition != nil
}

// Condition returns the condition, if any
func (obj *element) Condition() conditions.Condition {
	return obj.condition
}
