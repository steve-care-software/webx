package links

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/libraries/links/elements/conditions"
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
func NewElementWithConditionForTests(layer hash.Hash, condition conditions.Condition) Element {
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
