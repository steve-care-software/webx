package elements

import (
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/conditions"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers"
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
func NewElementWithConditionForTests(layer layers.Layer, condition conditions.Condition) Element {
	ins, err := NewElementBuilder().Create().WithLayer(layer).WithCondition(condition).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewElementForTests creates a new element for tests
func NewElementForTests(layer layers.Layer) Element {
	ins, err := NewElementBuilder().Create().WithLayer(layer).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
