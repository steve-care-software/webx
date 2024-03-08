package links

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/libraries/links/elements/conditions/resources"
	"github.com/steve-care-software/datastencil/domain/libraries/links/origins"
)

// NewLinksForTests creates a new links for tests
func NewLinksForTests(list []Link) Links {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewLinkForTests creates new link for tests
func NewLinkForTests(origin origins.Origin, elements Elements) Link {
	ins, err := NewLinkBuilder().Create().WithOrigin(origin).WithElements(elements).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewElementsForTests creates new elements for tests
func NewElementsForTests(list []Element) Elements {
	ins, err := NewElementsBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewElementWithConditionForTests creates a new element with condition  for tests
func NewElementWithConditionForTests(layer hash.Hash, condition Condition) Element {
	ins, err := NewElementBuilder().Create().WithLayer(layer).WithCondition(condition).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewElementForTests creates a new element for tests
func NewElementForTests(layer hash.Hash) Element {
	ins, err := NewElementBuilder().Create().WithLayer(layer).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewConditionWithNextForTests creates a new condition with next for tests
func NewConditionWithNextForTests(resource resources.Resource, next ConditionValue) Condition {
	ins, err := NewConditionBuilder().Create().WithResource(resource).WithNext(next).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewConditionForTests creates a new condition for tests
func NewConditionForTests(resource resources.Resource) Condition {
	ins, err := NewConditionBuilder().Create().WithResource(resource).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewConditionValueWithConditionForTests creates a new condition value with condition for tests
func NewConditionValueWithConditionForTests(condition Condition) ConditionValue {
	ins, err := NewConditionValueBuilder().Create().WithCondition(condition).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewConditionValueWithResourceForTests creates a new condition value with resource for tests
func NewConditionValueWithResourceForTests(resource resources.Resource) ConditionValue {
	ins, err := NewConditionValueBuilder().Create().WithResource(resource).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
