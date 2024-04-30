package conditions

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/queries/conditions/operators"
	"github.com/steve-care-software/datastencil/domain/instances/queries/conditions/pointers"
)

type condition struct {
	hash     hash.Hash
	pointer  pointers.Pointer
	operator operators.Operator
	element  Element
}

func createCondition(
	hash hash.Hash,
	pointer pointers.Pointer,
	operator operators.Operator,
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
func (obj *condition) Pointer() pointers.Pointer {
	return obj.pointer
}

// Operator returns the operator
func (obj *condition) Operator() operators.Operator {
	return obj.operator
}

// Element returns the element
func (obj *condition) Element() Element {
	return obj.element
}
