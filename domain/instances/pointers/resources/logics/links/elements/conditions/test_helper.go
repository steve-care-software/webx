package conditions

import "github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/links/elements/conditions/resources"

// NewConditionWithNextForTests creates a new condition with next for tests
func NewConditionWithNextForTests(resource resources.Resource, next Condition) Condition {
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
