package links

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/libraries/links/origins/operators"
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
func NewLinkForTests(origin Origin, elements Elements) Link {
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
func NewConditionWithNextForTests(resource ConditionResource, next ConditionValue) Condition {
	ins, err := NewConditionBuilder().Create().WithResource(resource).WithNext(next).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewConditionForTests creates a new condition for tests
func NewConditionForTests(resource ConditionResource) Condition {
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
func NewConditionValueWithResourceForTests(resource ConditionResource) ConditionValue {
	ins, err := NewConditionValueBuilder().Create().WithResource(resource).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewConditionResourceWithIsRaisedInLayerForTests creates a new condition resource withIsRaisedInLayer for tests
func NewConditionResourceWithIsRaisedInLayerForTests(code uint) ConditionResource {
	ins, err := NewConditionResourceBuilder().Create().WithCode(code).IsRaisedInLayer().Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewConditionResourceForTests creates a new condition resource for tests
func NewConditionResourceForTests(code uint) ConditionResource {
	ins, err := NewConditionResourceBuilder().Create().WithCode(code).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewOriginForTests creates a new origin for tests
func NewOriginForTests(resource OriginResource, operator operators.Operator, next OriginValue) Origin {
	ins, err := NewOriginBuilder().Create().WithResource(resource).WithOperator(operator).WithNext(next).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewOriginValueWithOriginForTests creates a new origin value with origin for tests
func NewOriginValueWithOriginForTests(origin Origin) OriginValue {
	ins, err := NewOriginValueBuilder().Create().WithOrigin(origin).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewOriginValueWithResourceForTests creates a new origin value with resource for tests
func NewOriginValueWithResourceForTests(resource OriginResource) OriginValue {
	ins, err := NewOriginValueBuilder().Create().WithResource(resource).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewOriginResourceWithIsMandatoryForTests creates a new origin resource with isMandatory for tests
func NewOriginResourceWithIsMandatoryForTests(layer hash.Hash) OriginResource {
	ins, err := NewOriginResourceBuilder().Create().WithLayer(layer).IsMandatory().Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewOriginResourceForTests creates a new origin resource for tests
func NewOriginResourceForTests(layer hash.Hash) OriginResource {
	ins, err := NewOriginResourceBuilder().Create().WithLayer(layer).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
