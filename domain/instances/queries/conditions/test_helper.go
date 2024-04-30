package conditions

import (
	"github.com/steve-care-software/datastencil/domain/instances/queries/conditions/operators"
	"github.com/steve-care-software/datastencil/domain/instances/queries/conditions/pointers"
	"github.com/steve-care-software/datastencil/domain/instances/queries/conditions/resources"
)

// NewConditionForTests creates a new condition for tests
func NewConditionForTests(pointer pointers.Pointer, operator operators.Operator, element Element) Condition {
	ins, err := NewBuilder().Create().
		WithPointer(pointer).
		WithOperator(operator).
		WithElement(element).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

// NewElementWithResourceForTests creates a new element with resource for tests
func NewElementWithResourceForTests(resource resources.Resource) Element {
	ins, err := NewElementBuilder().Create().WithResource(resource).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewElementWithConditionForTests creates a new element with condition for tests
func NewElementWithConditionForTests(condition Condition) Element {
	ins, err := NewElementBuilder().Create().WithCondition(condition).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
