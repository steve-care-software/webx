package conditions

import "github.com/steve-care-software/datastencil/domain/hash"

type condition struct {
	hash     hash.Hash
	pointer  Pointer
	operator Operator
	element  Element
}

func createCondition(
	hash hash.Hash,
	pointer Pointer,
	operator Operator,
	element Element,
) Condition {
	out := condition{
		hash:     hash,
		pointer:  pointer,
		operator: operator,
		element:  element,
	}

	return &out
}

// Hash returns the hash
func (obj *condition) Hash() hash.Hash {
	return obj.hash
}

// Pointer returns the pointer
func (obj *condition) Pointer() Pointer {
	return obj.pointer
}

// Operator returns the operator
func (obj *condition) Operator() Operator {
	return obj.operator
}

// Element returns the element
func (obj *condition) Element() Element {
	return obj.element
}
