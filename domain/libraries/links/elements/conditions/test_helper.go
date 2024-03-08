package conditions

import "github.com/steve-care-software/datastencil/domain/libraries/links/elements/conditions/resources"

// NewConditionWithNextForTests creates a new condition with next for tests
func NewConditionWithNextForTests(resource resources.Resource, next ConditionValue) Condition {
	ins, err := NewBuilder().Create().WithResource(resource).WithNext(next).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewConditionForTests creates a new condition for tests
func NewConditionForTests(resource resources.Resource) Condition {
	ins, err := NewBuilder().Create().WithResource(resource).Now()
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
