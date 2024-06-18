package conditions

import "github.com/steve-care-software/datastencil/domain/instances/pointers/conditions/operators"

// NewConditionForTests creates a new condition for tests
func NewConditionForTests(resource Resource) Condition {
	ins, err := NewBuilder().Create().WithResource(resource).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewConditionForTests creates a new condition for tests
func NewConditionWithComparisonsForTests(resource Resource, comparisons Comparisons) Condition {
	ins, err := NewBuilder().Create().WithResource(resource).WithComparisons(comparisons).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewResourceForTests creates a new resource for tests
func NewResourceForTests(path []string, mustBeLoaded bool) Resource {
	builder := NewResourceBuilder().Create().WithPath(path)
	if mustBeLoaded {
		builder.MustBeLoaded()
	}

	ins, err := builder.Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewComparisonsForTests creates new comparisons for tests
func NewComparisonsForTests(list []Comparison) Comparisons {
	ins, err := NewComparisonsBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewComparisonForTests creates a new comparison for tests
func NewComparisonForTests(operator operators.Operator, condition Condition) Comparison {
	ins, err := NewComparisonBuilder().Create().
		WithOperator(operator).
		WithCondition(condition).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
