package elements

import (
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/conditions"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/logics"
)

// NewElementsForTests creates new elements for tests
func NewElementsForTests(list []Element) Elements {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewElementWithConditionForTests creates a new element with condition  for tests
func NewElementWithConditionForTests(logic logics.Logic, condition conditions.Condition) Element {
	ins, err := NewElementBuilder().Create().WithLogic(logic).WithCondition(condition).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewElementForTests creates a new element for tests
func NewElementForTests(logic logics.Logic) Element {
	ins, err := NewElementBuilder().Create().WithLogic(logic).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
